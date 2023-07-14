package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler/response"
)

type UserUseCase interface {
	FindAllUsers(ctx context.Context, pagination request.Pagination) ([]response.User, error)
}
