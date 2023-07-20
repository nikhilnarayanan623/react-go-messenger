package response

import "time"

type Chat struct {
	ID             uint      `json:"chat_id" `
	UserID         uint      `json:"user_id"`
	FirstName      string    `json:"first_name"`
	UserName       string    `json:"user_name"`
	ProfilePicture string    `json:"profile_picture"`
	LastMessage    string    `json:"last_message"`
	LastMessageAt  time.Time `json:"last_message_at"`
}

type Message struct {
	ID            uint      `json:"message_id" `
	IsCurrentUser bool      `json:"is_current_user"`
	Content       string    `json:"content" gorm:"not null"`
	CreatedAt     time.Time `json:"created_at" gorm:"not null"`
}
