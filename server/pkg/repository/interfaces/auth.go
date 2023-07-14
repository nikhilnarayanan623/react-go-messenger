package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/domain"
)

type AuthRepository interface {
	SaveRefreshSession(ctx context.Context, refreshSession domain.RefreshSession) error
	FindRefreshSessionByTokenID(ctx context.Context, tokenID string) (domain.RefreshSession, error)
}
