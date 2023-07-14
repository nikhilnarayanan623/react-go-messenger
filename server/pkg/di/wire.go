//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/middleware"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/config"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/db"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/repository"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/service/google"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/service/token"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.Server, error) {

	wire.Build(
		db.ConnectDatabase,

		token.NewTokenService,
		google.NewGoogleAuth,
		middleware.NewMiddleware,

		repository.NewUserRepository,
		repository.NewAuthRepository,
		repository.NewChatRepository,

		usecase.NewAuthUseCase,
		usecase.NewChatUseCase,
		usecase.NewUserUseCase,

		handler.NewAuthHandler,
		handler.NewChatHandler,
		handler.NewUserHandler,

		http.NewServerHTTP,
	)

	return &http.Server{}, nil
}
