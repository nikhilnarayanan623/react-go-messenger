package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler/response"
)

type ChatRepository interface {
	FindAllRecentChatsOfUser(ctx context.Context, userID uint, pagination request.Pagination) ([]response.Chat, error)
	FindChatIDByUser1AndUser2ID(ctx context.Context, user1ID, user2ID uint) (chatID uint, err error)
	SaveChat(ctx context.Context, user1ID, user2ID uint) (chatID uint, err error)
}
