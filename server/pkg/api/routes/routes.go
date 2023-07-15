package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler/interfaces"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/middleware"
)

func SetupRoutes(api *gin.Engine, authHandler interfaces.AuthHandler,
	middleware middleware.Middleware, userHandler interfaces.UserHandler,
	chatHandler interfaces.ChatHandler) {

	{ // auth side
		api.POST(SignUpUrl, authHandler.UserSignUp)
		api.POST(SignInUrl, authHandler.UserLogin)
		api.POST(GoogleSignInUrl, authHandler.UserGoogleLogin)
		api.POST(RenewAccessTokenUrl, authHandler.UserRenewAccessToken())
	}

	api.Use(middleware.AuthenticateUser())

	{ // user
		api.POST(ListAllUsersUrl, userHandler.ListUsers)
	}
}
