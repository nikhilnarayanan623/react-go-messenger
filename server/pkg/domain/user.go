package domain

import "time"

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey;unique"`
	Age         uint      `json:"age" binding:"required,numeric"`
	GoogleImage string    `json:"google_profile_image"`
	FirstName   string    `json:"first_name" gorm:"not null" binding:"required,min=2,max=50"`
	LastName    string    `json:"last_name" gorm:"not null" binding:"required,min=1,max=50"`
	UserName    string    `json:"user_name" gorm:"not null;unique" binding:"required,min=3,max=15"`
	Email       string    `json:"email" gorm:"unique;not null" binding:"required,email"`
	Password    string    `json:"password" binding:"required"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt   time.Time `json:"updated_at"`
}
