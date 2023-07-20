//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/middleware"
	socket "github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/service"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/config"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/db"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/repository"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/service/google"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/service/token"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/usecase"

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

		socket.NewWebSocketService,
		handler.NewAuthHandler,
		handler.NewChatHandler,
		handler.NewUserHandler,

		http.NewServerHTTP,
	)

	return &http.Server{}, nil
}
