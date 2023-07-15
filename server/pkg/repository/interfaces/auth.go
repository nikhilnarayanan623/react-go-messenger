package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/domain"
)

type AuthRepository interface {
	SaveRefreshSession(ctx context.Context, refreshSession domain.RefreshSession) error
	FindRefreshSessionByTokenID(ctx context.Context, tokenID string) (domain.RefreshSession, error)
}
