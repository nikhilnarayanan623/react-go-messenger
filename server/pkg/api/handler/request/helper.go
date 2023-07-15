package request

import "github.com/gin-gonic/gin"

func GetUserIdFromContext(ctx *gin.Context) uint {
	userID := ctx.GetUint("userId")
	return userID
}
