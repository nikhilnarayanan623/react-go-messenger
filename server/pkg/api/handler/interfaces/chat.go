package interfaces

import "github.com/gin-gonic/gin"

type ChatHandler interface {
	GetRecentChats(ctx *gin.Context)
	SaveChat(ctx *gin.Context)
	GetAllMessages(ctx *gin.Context)
	SaveMessage(ctx *gin.Context)
}
