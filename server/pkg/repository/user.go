package repository

import (
	"context"
	"time"

	"github.com/nikhilnarayanan623/go-socket-chat/server/pkg/domain"
	"github.com/nikhilnarayanan623/go-socket-chat/server/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{
		DB: DB,
	}
}

func (c *userDatabase) FindUserByUserID(ctx context.Context, userID uint) (user domain.User, err error) {

	query := `SELECT * FROM users WHERE id = $1`
	err = c.DB.Raw(query, userID).Scan(&user).Error

	return user, err
}

func (c *userDatabase) FindUserByEmail(ctx context.Context, email string) (user domain.User, err error) {

	query := `SELECT * FROM users WHERE email = $1`
	err = c.DB.Raw(query, email).Scan(&user).Error

	return user, err
}

func (c *userDatabase) FindUserByUserName(ctx context.Context, userName string) (user domain.User, err error) {

	query := `SELECT * FROM users WHERE user_name = $1`
	err = c.DB.Raw(query, userName).Scan(&user).Error

	return user, err
}

func (c *userDatabase) FindUserByUserNameEmailNotID(ctx context.Context,
	userDetails domain.User) (user domain.User, err error) {

	query := `SELECT * FROM users WHERE (user_name = $1 OR email = $2 ) AND id != $3`
	err = c.DB.Raw(query, userDetails.UserName, userDetails.Email, userDetails.ID).Scan(&user).Error

	return
}

func (c *userDatabase) SaveUser(ctx context.Context, user domain.User) (userID uint, err error) {

	query := `INSERT INTO users (user_name, first_name, 
		last_name, age, email, password, google_image, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8 ) RETURNING id`

	createdAt := time.Now()
	err = c.DB.Raw(query, user.UserName, user.FirstName, user.LastName,
		user.Age, user.Email, user.Password, user.GoogleImage, createdAt).Scan(&user).Error

	return userID, err
}
