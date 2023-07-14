package domain

import "time"

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey;unique"`
	Age         uint      `json:"age"`
	GoogleImage string    `json:"google_profile_image"`
	FirstName   string    `json:"first_name" gorm:"not null"`
	LastName    string    `json:"last_name" gorm:"not null"`
	UserName    string    `json:"user_name" gorm:"not null;unique"`
	Email       string    `json:"email" gorm:"unique;not null"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at"`
}
