package interfaces

import "github.com/gin-gonic/gin"

type UserHandler interface {
	ListAllUsers(ctx *gin.Context)
}
