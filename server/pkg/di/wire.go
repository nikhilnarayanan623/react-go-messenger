//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/nikhilnarayanan623/go-socket-chat/server/pkg/api"
	"github.com/nikhilnarayanan623/go-socket-chat/server/pkg/api/handler"
	"github.com/nikhilnarayanan623/go-socket-chat/server/pkg/config"
	"github.com/nikhilnarayanan623/go-socket-chat/server/pkg/db"
	"github.com/nikhilnarayanan623/go-socket-chat/server/pkg/repository"
	"github.com/nikhilnarayanan623/go-socket-chat/server/pkg/api/middleware"
	"github.com/nikhilnarayanan623/go-socket-chat/server/pkg/service/token"
	"github.com/nikhilnarayanan623/go-socket-chat/server/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.Server, error) {

	wire.Build(
		db.ConnectDatabase,

		token.NewTokenService,
		middleware.NewMiddleware,

		repository.NewUserRepository,
		repository.NewAuthRepository,
		repository.NewChatRepository,

		usecase.NewAuthUseCase,
		usecase.NewChatUseCase,

		handler.NewAuthHandler,
		handler.NewChatHandler,

		http.NewServerHTTP,
	)

	return &http.Server{}, nil
}
