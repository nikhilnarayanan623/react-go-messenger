package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler/interfaces"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/middleware"
)

func SetupRoutes(api *gin.Engine, authHandler interfaces.AuthHandler,
	middleware middleware.Middleware, userHandler interfaces.UserHandler,
	chatHandler interfaces.ChatHandler) {

	{ // auth
		api.POST(SignUpUrl, authHandler.UserSignUp)
		api.POST(SignInUrl, authHandler.UserLogin)
		api.POST(GoogleSignInUrl, authHandler.UserGoogleLogin)
		api.POST(RenewAccessTokenUrl, authHandler.UserRenewAccessToken())
	}

	// r := api.Use(middleware.AuthenticateUser())

	{ // user
		api.GET(ListAllUsersUrl, userHandler.ListUsers)
	}

	{ //chats

		api.GET(RecentChatsUrl, chatHandler.GetRecentChats)
		api.POST(CreateChatUrl, chatHandler.SaveChat)

		api.GET(ListAllMessagesUrl, chatHandler.GetAllMessages)
	}
}
