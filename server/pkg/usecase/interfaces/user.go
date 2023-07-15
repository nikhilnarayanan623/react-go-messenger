package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler/response"
)

type UserUseCase interface {
	FindAllUsers(ctx context.Context, pagination request.Pagination) ([]response.User, error)
}
