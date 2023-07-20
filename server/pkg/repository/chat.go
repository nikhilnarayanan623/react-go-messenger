package repository

import (
	"context"
	"time"

	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler/response"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/domain"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/repository/interfaces"

	"gorm.io/gorm"
)

type chatDatabase struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) interfaces.ChatRepository {
	return &chatDatabase{
		db: db,
	}
}

func (c *chatDatabase) FindAllRecentChatsOfUser(ctx context.Context, userID uint,
	pagination request.Pagination) (chats []response.Chat, err error) {

	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit

	query := `SELECT DISTINCT ON (chats.id) chats.id, he.first_name, he.user_name,
	he.google_image AS profile_picture, messages.content AS last_message, messages.created_at AS last_message_at
	FROM chats
	INNER JOIN users me ON  me.id = chats.user1_id OR me.id = chats.user2_id
	INNER JOIN users he ON he.id = chats.user1_id OR he.id = chats.user2_id
	INNER JOIN messages ON messages.chat_id = chats.id
	WHERE me.id = $1 AND he.id != $1 ORDER BY chats.id, messages.created_at DESC 
	LIMIT $2 OFFSET $3`

	err = c.db.Raw(query, userID, limit, offset).Scan(&chats).Error

	return
}

func (c *chatDatabase) SaveChat(ctx context.Context, user1ID, user2ID uint) (chatID uint, err error) {

	createdAt := time.Now()

	query := `INSERT INTO chats (user1_id, user2_id, created_at) VALUES($1, $2, $3) RETURNING id`

	err = c.db.Raw(query, user1ID, user2ID, createdAt).Scan(&chatID).Error

	return

}

func (c *chatDatabase) FindChatIDByUser1AndUser2ID(ctx context.Context, user1ID, user2ID uint) (chatID uint, err error) {

	query := `SELECT id FROM chats WHERE user1_id = $1 AND user2_id = $2`
	err = c.db.Raw(query, user1ID, user2ID).Scan(&chatID).Error

	return
}

func (c *chatDatabase) FindAllMessagesByChatAndUserID(ctx context.Context,
	chatID, userID uint, pagination request.Pagination) (messages []response.Message, err error) {

	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit

	query := `SELECT id, content, created_at, 
	CASE WHEN sender_id = $1 THEN 'T' ELSE 'F' END AS is_current_user FROM messages 
	WHERE chat_id = $2
	ORDER BY created_at DESC LIMIT $3 OFFSET $4`

	err = c.db.Raw(query, userID, chatID, limit, offset).Scan(&messages).Error

	return
}

func (c *chatDatabase) SaveMessage(ctx context.Context, message domain.Message) error {

	createdAt := time.Now()
	query := `INSERT INTO messages (chat_id, sender_id, content, created_at) VALUES($1, $2, $3, $4)`
	err := c.db.Exec(query, message.ChatID, message.SenderID, message.Content, createdAt).Error

	return err
}
func (c *chatDatabase) FindReceiverOfChatBySenderID(ctx context.Context,
	chatID, senderID uint) (receiverID uint, err error) {

	query := `SELECT CASE WHEN user1_id = $1 THEN user2_id ELSE user1_id END AS receiver_id FROM chats 
	WHERE id = $2`
	err = c.db.Raw(query, senderID, chatID).Scan(&receiverID).Error

	return
}
