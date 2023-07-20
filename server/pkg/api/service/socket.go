package socket

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/service/token"
)

type WebSocketService interface {
	ServeWebSocket(ctx *gin.Context)
}

type webSocketService struct {
	upgrader     websocket.Upgrader
	connections  map[uint]*websocket.Conn
	mu           sync.Mutex
	tokenService token.TokenService
}

type Message struct {
	RecieverID uint   `json:"reciever_id"`
	Content    string `json:"content"`
}

func NewWebSocketService(tokenService token.TokenService) WebSocketService {

	return &webSocketService{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
		connections:  make(map[uint]*websocket.Conn),
		mu:           sync.Mutex{},
		tokenService: tokenService,
	}
}

// ServeWebSocket godoc
// @Summary Sever Socket Connection (User)
// @Description API for user to create a web socket connection
// @Security ApiKeyAuth
// @Id ServeWebSocket
// @Tags Users Socket
// @Router /ws [get]
func (c *webSocketService) ServeWebSocket(ctx *gin.Context) {
	soketConn, err := c.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		log.Println("failed to upgrade connection")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	closeHand := soketConn.CloseHandler()

	userID, err := c.verifyConnection(ctx, soketConn)
	if err != nil {
		log.Println("failed to verify token", err)
		closeHand(websocket.ClosePolicyViolation, err.Error())
		return
	}

	c.mu.Lock()
	c.connections[userID] = soketConn
	c.mu.Unlock()

	{ // wait until the connectin close
		fmt.Println("successfully connected for user_id : ", userID)
		go c.readMessages(ctx, soketConn) // read messages
		<-ctx.Done()
		log.Println("connection closed from request")
	}

	c.mu.Lock()
	delete(c.connections, userID)
	c.mu.Unlock()
}

type TokenRequest struct {
	Token string `json:"token"`
}

func (c *webSocketService) verifyConnection(ctx context.Context, sc *websocket.Conn) (userID uint, err error) {

	tokenChan := make(chan TokenRequest)
	errChan := make(chan error)

	// wait for the token to send within 5 second of connectin established
	go func() {
		var body TokenRequest
		err := sc.ReadJSON(&body)
		if err != nil {
			errChan <- err
		}
		tokenChan <- body
	}()

	select {
	case err := <-errChan:
		return 0, err
	case body := <-tokenChan:
		tokenRes, err := c.tokenService.VerifyToken(token.VerifyTokenRequest{TokenString: body.Token, UsedFor: token.User})
		if err != nil {
			return 0, err
		}
		return tokenRes.UserID, nil
	case <-time.After(5 * time.Second):
		return 0, errors.New("time exceed for waiting token send")
	}

}

func (c *webSocketService) readMessages(ctx context.Context, sc *websocket.Conn) {

	messageChan := make(chan Message)
	go func() {
		for {
			var body Message
			if err := sc.ReadJSON(&body); err != nil {
				log.Println("failed to read message: ", err)
				return
			}
			messageChan <- body
		}

	}()

	// get each messge and if the whole connectin lost then return
	for {
		select {
		case message := <-messageChan:
			go c.sendMessge(message)
		case <-ctx.Done():
			return
		}
	}
}

func (c *webSocketService) sendMessge(messge Message) (received bool, err error) {

	c.mu.Lock()
	defer c.mu.Unlock()
	// find other connection
	conn, ok := c.connections[messge.RecieverID]
	if !ok {
		return false, nil
	}

	// send to other connection
	if err := conn.WriteJSON(messge); err != nil {
		log.Println("failed to write message: ", messge)
		return false, err
	}
	return true, nil
}
