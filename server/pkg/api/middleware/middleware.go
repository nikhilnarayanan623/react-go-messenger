package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/go-socket-chat/server/pkg/service/token"
)

type Middleware interface {
	AuthenticateUser() gin.HandlerFunc
}

type middleware struct {
	tokenService token.TokenService
}

func NewMiddleware(tokenService token.TokenService) Middleware {
	return &middleware{
		tokenService: tokenService,
	}
}
