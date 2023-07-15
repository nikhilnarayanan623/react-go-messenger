package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler/response"
	repoInterface "github.com/nikhilnarayanan623/server/react-go-messenger/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/utils"
)

type chatUseCase struct {
	chatRepo repoInterface.ChatRepository
}

func NewChatUseCase(chatRepo repoInterface.ChatRepository) interfaces.ChatUseCase {
	return &chatUseCase{
		chatRepo: chatRepo,
	}
}

func (c *chatUseCase) FindAllRecentChatsOfUser(ctx *gin.Context, userID uint, pagination request.Pagination) ([]response.Chat, error) {

	chats, err := c.chatRepo.FindAllRecentChatsOfUser(ctx, userID, pagination)
	if err != nil {
		return nil, utils.PrependMessageToError(err, "failed to get user chats from database")
	}

	return chats, nil
}
