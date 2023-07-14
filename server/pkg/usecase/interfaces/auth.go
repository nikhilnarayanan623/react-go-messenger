package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/domain"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/service/token"
)

type AuthUseCase interface {
	//user
	UserSignUp(ctx context.Context, signUpDetails domain.User) error
	UserLogin(ctx context.Context, loginDetails request.Login) (userID uint, err error)
	GoogleLogin(ctx context.Context, token string) (userID uint, err error)
	// token
	GenerateAccessToken(ctx context.Context, tokenParams GenerateTokenParams) (tokenString string, err error)
	GenerateRefreshToken(ctx context.Context, tokenParams GenerateTokenParams) (tokenString string, err error)
	VerifyAndGetRefreshTokenSession(ctx context.Context, refreshToken string, usedFor token.UserType) (domain.RefreshSession, error)
}

type GenerateTokenParams struct {
	UserID   uint
	UserType token.UserType
}
