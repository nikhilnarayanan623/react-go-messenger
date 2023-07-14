package repository

import (
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
