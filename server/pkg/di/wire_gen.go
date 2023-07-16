// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler/service"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/middleware"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/config"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/db"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/repository"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/service/google"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/service/token"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.Server, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	authRepository := repository.NewAuthRepository(gormDB)
	tokenService := token.NewTokenService(cfg)
	googleAuth := google.NewGoogleAuth(cfg)
	userRepository := repository.NewUserRepository(gormDB)
	authUseCase := usecase.NewAuthUseCase(authRepository, tokenService, googleAuth, userRepository)
	authHandler := handler.NewAuthHandler(authUseCase)
	middlewareMiddleware := middleware.NewMiddleware(tokenService)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	chatRepository := repository.NewChatRepository(gormDB)
	chatUseCase := usecase.NewChatUseCase(chatRepository)
	socketService := service.NewSocketService(tokenService)
	chatHandler := handler.NewChatHandler(chatUseCase, socketService)
	server := http.NewServerHTTP(cfg, authHandler, middlewareMiddleware, userHandler, chatHandler)
	return server, nil
}
