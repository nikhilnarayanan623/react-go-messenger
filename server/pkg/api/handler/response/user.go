package response

import "time"

type User struct {
	ID           uint      `json:"user_id"`
	ProfileImage string    `json:"profile_image"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	UserName     string    `json:"user_name"`
	CreatedAt    time.Time `json:"created_at"`
}
