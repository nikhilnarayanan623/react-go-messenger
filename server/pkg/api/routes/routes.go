package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler/interfaces"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/middleware"
	socket "github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/service"
)

func SetupRoutes(api *gin.Engine, authHandler interfaces.AuthHandler,
	middleware middleware.Middleware, userHandler interfaces.UserHandler,
	chatHandler interfaces.ChatHandler, socketService socket.WebSocketService) {

	auth := api.Group("")
	{ // auth
		auth.POST(SignUpUrl, authHandler.UserSignUp)
		auth.POST(SignInUrl, authHandler.UserLogin)
		auth.POST(GoogleSignInUrl, authHandler.UserGoogleLogin)
		auth.POST(RenewAccessTokenUrl, authHandler.UserRenewAccessToken())
	}

	api.GET(SocketUrl, socketService.ServeWebSocket)

	authorized := api.Group("", middleware.AuthenticateUser())

	{ // user
		authorized.GET(ListAllUsersUrl, userHandler.ListUsers)
	}

	{ //chats

		authorized.GET(RecentChatsUrl, chatHandler.GetRecentChats)
		authorized.POST(CreateChatUrl, chatHandler.SaveChat)

		authorized.GET(ListAllMessagesUrl, chatHandler.GetAllMessages)
		authorized.POST(SaveMessageUrl, chatHandler.SaveMessage)
	}
}
