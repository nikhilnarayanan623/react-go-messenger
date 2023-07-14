package repository

import (
	"context"
	"time"

	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler/response"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/domain"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type userDatabase struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userDatabase{
		db: db,
	}
}

func (c *userDatabase) FindUserByUserID(ctx context.Context, userID uint) (user domain.User, err error) {

	query := `SELECT * FROM users WHERE id = $1`
	err = c.db.Raw(query, userID).Scan(&user).Error

	return user, err
}

func (c *userDatabase) FindUserByEmail(ctx context.Context, email string) (user domain.User, err error) {

	query := `SELECT * FROM users WHERE email = $1`
	err = c.db.Raw(query, email).Scan(&user).Error

	return user, err
}

func (c *userDatabase) FindUserByUserName(ctx context.Context, userName string) (user domain.User, err error) {

	query := `SELECT * FROM users WHERE user_name = $1`
	err = c.db.Raw(query, userName).Scan(&user).Error

	return user, err
}

func (c *userDatabase) FindUserByUserNameEmailNotID(ctx context.Context,
	userDetails domain.User) (user domain.User, err error) {

	query := `SELECT * FROM users WHERE (user_name = $1 OR email = $2 ) AND id != $3`
	err = c.db.Raw(query, userDetails.UserName, userDetails.Email, userDetails.ID).Scan(&user).Error

	return
}

func (c *userDatabase) SaveUser(ctx context.Context, user domain.User) (userID uint, err error) {

	query := `INSERT INTO users (user_name, first_name, 
		last_name, age, email, password, google_image, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8 ) RETURNING id`

	createdAt := time.Now()
	err = c.db.Raw(query, user.UserName, user.FirstName, user.LastName,
		user.Age, user.Email, user.Password, user.GoogleImage, createdAt).Scan(&user).Error

	return userID, err
}

func (c *userDatabase) FindAllUsers(ctx context.Context,
	pagination request.Pagination) (users []response.User, err error) {

	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit

	query := `SELECT id, first_name, last_name, user_name, google_image AS profile_image, created_at 
	FROM users LIMIT $1 OFFSET $2`

	err = c.db.Raw(query, limit, offset).Scan(&users).Error

	return
}
