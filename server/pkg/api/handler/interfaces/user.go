package interfaces

import "github.com/gin-gonic/gin"

type UserHandler interface {
	ListUsers(ctx *gin.Context)
}
