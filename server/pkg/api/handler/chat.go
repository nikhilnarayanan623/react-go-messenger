package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler/interfaces"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler/response"
	usecase "github.com/nikhilnarayanan623/server/react-go-messenger/pkg/usecase/interfaces"
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
// @Router /api/chats [get]
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
