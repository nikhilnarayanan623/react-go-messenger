package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler/response"
)

type ChatUseCase interface {
	FindAllRecentChatsOfUser(ctx *gin.Context, userID uint, pagination request.Pagination) ([]response.Chat, error)
}
