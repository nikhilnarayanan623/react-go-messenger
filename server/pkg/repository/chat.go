package repository

import (
	"context"
	"time"

	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler/response"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/repository/interfaces"

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
