package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler/response"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/service/token"
)

const (
	authorizationHeaderKey string = "authorization"
	authorizationType      string = "bearer"
)

func (c *middleware) AuthenticateUser() gin.HandlerFunc {
	// return c.authorize(token.User)
	return c.authorize(token.User)
}

func (c *middleware) authorize(tokenUser token.UserType) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authorizationValues := ctx.GetHeader(authorizationHeaderKey)

		authFields := strings.Fields(authorizationValues)
		if len(authFields) < 2 {
			c.middlewareUsingCookie(ctx, tokenUser)

			// err := errors.New("authorization token not provided")

			// response.ErrorResponse(ctx, http.StatusUnauthorized, "Failed to authorize request", err, nil)
			// ctx.Abort()
			return
		}

		authType := authFields[0]
		accessToken := authFields[1]

		if !strings.EqualFold(authType, authorizationType) {
			err := errors.New("invalid authorization type")
			response.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized user", err, nil)
			ctx.Abort()
			return
		}

		tokenVerifyReq := token.VerifyTokenRequest{
			TokenString: accessToken,
			UsedFor:     tokenUser,
		}

		verifyRes, err := c.tokenService.VerifyToken(tokenVerifyReq)

		if err != nil {
			response.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized user", err, nil)
			ctx.Abort()
			return
		}

		ctx.Set("userId", verifyRes.UserID)
	}
}

// for swagger
func (c *middleware) middlewareUsingCookie(ctx *gin.Context, tokenUser token.UserType) {

	cookieName := "auth-" + string(tokenUser)
	accessToken, err := ctx.Cookie(cookieName)

	if err != nil || accessToken == "" {
		if accessToken == "" {
			err = errors.Join(err, errors.New("access token not found"))
		}
		response.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized user", err, nil)
		ctx.Abort()
		return
	}
	verifyRes, err := c.tokenService.VerifyToken(token.VerifyTokenRequest{
		TokenString: accessToken,
		UsedFor:     tokenUser,
	})

	if err != nil {
		response.ErrorResponse(ctx, http.StatusUnauthorized, "Unauthorized user", err, nil)
		ctx.Abort()
		return
	}

	ctx.Set("userId", verifyRes.UserID)
}
