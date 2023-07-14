package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/go-socket-chat/server/pkg/domain"
)

type AuthRepository interface {
	SaveRefreshSession(ctx context.Context, refreshSession domain.RefreshSession) error
	FindRefreshSessionByTokenID(ctx context.Context, tokenID string) (domain.RefreshSession, error)
}
