package service

import (
	"context"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/service/token"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/utils"
)

type SocketService interface {
	Upgrade(ctx context.Context, w http.ResponseWriter, r *http.Request)
	SendMessage(userID uint, message Message) (isReceived bool)
}

type socketService struct {
	upgrader     websocket.Upgrader
	Connections  map[uint]*websocket.Conn
	mu           sync.Mutex
	tokenService token.TokenService
}
type Message struct {
	ChatID  uint   `json:"chat_id"`
	Content string `json:"content"`
}
type Token struct {
	Token string `json:"token"`
}

func NewSocketService(tokenService token.TokenService) SocketService {

	return &socketService{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			}, // for all origins
		},
		Connections:  make(map[uint]*websocket.Conn), // to store all connections
		mu:           sync.Mutex{},
		tokenService: tokenService,
	}
}

func (c *socketService) Upgrade(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	socketConn, err := c.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	closeHandler := socketConn.CloseHandler()

	// validate the token from connect
	var body Token
	if err := socketConn.ReadJSON(&body); err != nil {

		err := utils.PrependMessageToError(err, "failed to get token from connection message")

		closeHandler(websocket.ClosePolicyViolation, err.Error())
	}

	tokenRes, err := c.tokenService.VerifyToken(token.VerifyTokenRequest{
		TokenString: body.Token,
		UsedFor:     token.User,
	})

	if err != nil {
		err = utils.PrependMessageToError(err, "failed to verify token")
		closeHandler(websocket.ClosePolicyViolation, err.Error())
	}

	// token is valid then store the connections
	c.mu.Lock()
	c.Connections[tokenRes.UserID] = socketConn // add new socket connection on connections
	c.mu.Unlock()

	socketConn.WriteJSON("successfully connection upgraded")

	<-ctx.Done() // wait util connection close

	// last remove the socket connection
	closeHandler(websocket.CloseGoingAway, "connection closed from request")
	delete(c.Connections, tokenRes.UserID)
}

func (c *socketService) SendMessage(userID uint, message Message) bool {

	c.mu.Lock()
	defer c.mu.Unlock()

	receiverConn, ok := c.Connections[userID]
	if !ok { // check receiver is online or not
		return false
	}

	if err := receiverConn.WriteJSON(message); err != nil {
		return false
	}

	return true
}
