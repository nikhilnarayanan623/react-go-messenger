package response

import "time"

type Chat struct {
	ID             uint      `json:"chat_id" `
	FirstName      string    `json:"first_name"`
	UserName       string    `json:"user_name"`
	ProfilePicture string    `json:"profile_picture"`
	LastMessage    string    `json:"last_message"`
	LastMessageAt  time.Time `json:"last_message_at"`
}
