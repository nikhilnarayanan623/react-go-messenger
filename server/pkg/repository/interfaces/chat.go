package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler/response"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/domain"
)

type ChatRepository interface {
	FindAllRecentChatsOfUser(ctx context.Context, userID uint, pagination request.Pagination) ([]response.Chat, error)
	FindChatIDByUser1AndUser2ID(ctx context.Context, user1ID, user2ID uint) (chatID uint, err error)
	SaveChat(ctx context.Context, user1ID, user2ID uint) (chatID uint, err error)
	FindAllMessagesByChatAndUserID(ctx context.Context, chatID, userID uint,
		pagination request.Pagination) ([]response.Message, error)
	SaveMessage(ctx context.Context, message domain.Message) error
	FindReceiverOfChatBySenderID(ctx context.Context, chatID, senderID uint) (receiverID uint, err error)
}
