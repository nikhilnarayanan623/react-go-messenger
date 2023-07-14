package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler/interfaces"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/middleware"
)

func SetupRoutes(api *gin.RouterGroup, authHandler interfaces.AuthHandler,
	middleware middleware.Middleware, chatHandler interfaces.ChatHandler) {

	auth := api.Group("/auth")
	{
		signup := auth.Group("/signup")
		{
			signup.POST("/", authHandler.UserSignUp)
		}

		login := auth.Group("/login")
		{
			login.POST("/", authHandler.UserLogin)
		}

		auth.POST("/google-auth", authHandler.UserGoogleLogin)

		auth.POST("/renew-access-token", authHandler.UserRenewAccessToken())

		// api.POST("/logout")

	}

}
