package usecase

import (
	"context"

	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler/response"
	repo "github.com/nikhilnarayanan623/react-go-messenger/server/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/utils"
)

type userUseCase struct {
	repo repo.UserRepository
}

func NewUserUseCase(repo repo.UserRepository) interfaces.UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}

func (c *userUseCase) FindAllUsers(ctx context.Context, pagination request.Pagination) ([]response.User, error) {

	users, err := c.repo.FindAllUsers(ctx, pagination)
	if err != nil {
		return nil, utils.PrependMessageToError(err, "failed to get all users from database")
	}

	return users, nil
}
