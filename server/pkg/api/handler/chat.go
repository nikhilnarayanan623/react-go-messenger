package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler/interfaces"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler/response"
	usecase "github.com/nikhilnarayanan623/react-go-messenger/server/pkg/usecase/interfaces"
)

type chatHandler struct {
	usecase usecase.ChatUseCase
}

func NewChatHandler(usecase usecase.ChatUseCase) interfaces.ChatHandler {
	return &chatHandler{
		usecase: usecase,
	}
}

// GetRecentChats godoc
// @Summary Get user chats (User)
// @Description API for user to get all recent chats of user with others
// @Security ApiKeyAuth
// @Id GetRecentChats
// @Tags Users Chats
// @Param page_number query int false "Page Number"
// @Param count query int false "Count"
// @Router /chats [get]
// @Success 200 {object} response.Response{data=[]response.Chat} "Successfully retrieved recent chats of user"
// @Success 204 {object} response.Response{} "There is no chats recent chats for users"
// @Failure 500 {object} response.Response{} "Failed to retrieved recent chats of user"
func (c *chatHandler) GetRecentChats(ctx *gin.Context) {

	userID := request.GetUserIdFromContext(ctx)
	pagination := request.GetPagination(ctx)

	chats, err := c.usecase.FindAllRecentChatsOfUser(ctx, userID, pagination)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieved recent chats of user", err, nil)
		return
	}

	if len(chats) == 0 {
		response.SuccessResponse(ctx, http.StatusNoContent, "There is no chats recent chats for users", nil)
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, "Successfully retrieved recent chats of user", chats)
}

// SaveChat godoc
// @Summary Save New chat (User)
// @Description API for user create a new chat with other user if already exist will return then existing chat id
// @Security ApiKeyAuth
// @Id SaveChat
// @Tags Users Chats
// @Param input body request.Chat true "Input fields"
// @Router /chats [post]
// @Success 200 {object} response.Response{data=uint} "Successfully chat saved"
// @Failure 500 {object} response.Response{} "Failed save to chat for user"
func (c *chatHandler) SaveChat(ctx *gin.Context) {

	var body request.Chat

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, body)
		return
	}
	userID := request.GetUserIdFromContext(ctx)

	chatID, err := c.usecase.SaveChat(ctx, userID, body.OtherUserID)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to chat for user", err, nil)
		return
	}

	response.SuccessResponse(ctx, http.StatusCreated, "Successfully chat saved", chatID)
}

// GetAllMessages godoc
// @Summary Get chats messages (User)
// @Description API for user to get all messages in a specific chat
// @Security ApiKeyAuth
// @Id GetAllMessages
// @Tags Users Chats
// @Param chat_id path int true "Chat ID"
// @Param page_number query int false "Page Number"
// @Param count query int false "Count"
// @Router /chats/{chat_id}/messages [get]
// @Success 200 {object} response.Response{data=[]response.Message} "Successfully retrieved message for the chat"
// @Success 204 {object} response.Response{} "There is no message between users"
// @Failure 500 {object} response.Response{} "Failed to retrieve message for this chat"
func (c *chatHandler) GetAllMessages(ctx *gin.Context) {

	chatID, err := request.GetParamAsUint(ctx, "chat_id")
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindParamFailMessage, err, nil)
		return
	}
	pagination := request.GetPagination(ctx)
	userID := request.GetUserIdFromContext(ctx)

	messages, err := c.usecase.FindAllMessagesOfUserForAChat(ctx, chatID, userID, pagination)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve message for this chat", err, nil)
	}

	if len(messages) == 0 {
		response.SuccessResponse(ctx, http.StatusNoContent, "There is no message between users", err)
	}

	response.SuccessResponse(ctx, http.StatusOK, "Successfully retrieved message for the chat", messages)
}
