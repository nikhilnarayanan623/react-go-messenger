package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler/interfaces"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler/response"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/domain"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/service/token"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/usecase"
	usecaseInterface "github.com/nikhilnarayanan623/server/react-go-messenger/pkg/usecase/interfaces"
)

type AuthHandler struct {
	authUseCase usecaseInterface.AuthUseCase
}

func NewAuthHandler(authUsecase usecaseInterface.AuthUseCase) interfaces.AuthHandler {
	return &AuthHandler{
		authUseCase: authUsecase,
	}
}

// UserLogin godoc
// @Summary Login with password (User)
// @Description API for user to login with email | phone | user_name with password
// @Security ApiKeyAuth
// @Id UserLogin
// @Tags User Authentication
// @Param        inputs   body     request.Login{}   true  "Login Details"
// @Router /api/auth/login [post]
// @Success 200 {object} response.Response{data=response.TokenResponse{}} "Successfully logged in"
// @Failure 400 {object} response.Response{}  "Invalid inputs"
// @Failure 403 {object} response.Response{}  "User blocked by admin"
// @Failure 401 {object} response.Response{}  "User not exist with given login credentials"
// @Failure 500 {object} response.Response{}  "Failed to login"
func (c *AuthHandler) UserLogin(ctx *gin.Context) {

	var body request.Login

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, body)
		return
	}

	userID, err := c.authUseCase.UserLogin(ctx, body)

	if err != nil {

		var statusCode int

		switch {
		case errors.Is(err, usecase.ErrEmptyLoginCredentials):
			statusCode = http.StatusBadRequest
		case errors.Is(err, usecase.ErrUserNotExist):
			statusCode = http.StatusNotFound
		case errors.Is(err, usecase.ErrWrongPassword):
			statusCode = http.StatusUnauthorized
		default:
			statusCode = http.StatusInternalServerError
		}

		response.ErrorResponse(ctx, statusCode, "Failed to login", err, nil)
		return
	}

	// common functionality for admin and user
	c.setupTokenAndResponse(ctx, token.User, userID)
}

// UserGoogleLogin godoc
// @Summary Login with google (User)
// @Description API for user to login with google
// @Security ApiKeyAuth
// @Id UserGoogleLogin
// @Tags User Authentication
// @Param        inputs   body     request.GoogleLogin{}   true  "Google Token Input"
// @Router /api/auth/google-auth [post]
// @Success 200 {object} response.Response{data=response.TokenResponse{}} "Successfully logged in with google"
// @Failure 400 {object} response.Response{}  "Invalid inputs"
// @Failure 500 {object} response.Response{}  "Failed to login"
func (c *AuthHandler) UserGoogleLogin(ctx *gin.Context) {
	var body request.GoogleLogin

	if err := ctx.ShouldBind(&body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, body)
		return
	}

	userID, err := c.authUseCase.GoogleLogin(ctx, body.TokenCode)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to login with google", err, nil)
		return
	}

	c.setupTokenAndResponse(ctx, token.User, userID)
}

func (c *AuthHandler) setupTokenAndResponse(ctx *gin.Context, tokenUser token.UserType, userID uint) {

	tokenParams := usecaseInterface.GenerateTokenParams{
		UserID:   userID,
		UserType: tokenUser,
	}

	accessToken, err := c.authUseCase.GenerateAccessToken(ctx, tokenParams)

	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to generate access token", err, nil)
		return
	}

	refreshToken, err := c.authUseCase.GenerateRefreshToken(ctx, usecaseInterface.GenerateTokenParams{
		UserID:   userID,
		UserType: tokenUser,
	})
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to generate refresh token", err, nil)
		return
	}

	// authorizationValue := authorizationType + " " + accessToken
	// ctx.Header(authorizationHeaderKey, authorizationValue)

	//ctx.Header("access_token", accessToken)
	//ctx.Header("refresh_token", refreshToken)

	{ // for swagger documentation set up cookie
		cookieName := "auth-" + string(tokenUser)
		ctx.SetCookie(cookieName, accessToken, 15*60, "", "", false, true)
	}

	tokenRes := response.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	response.SuccessResponse(ctx, http.StatusOK, "Successfully logged in", tokenRes)
}

// UserSignUp godoc
// @Summary Signup (User)
// @Description API for user to register a new account
// @Security ApiKeyAuth
// @Id UserSignUp
// @Tags User Authentication
// @Param input body request.UserSignUp{} true "Input Fields"
// @Router /api/auth/signup [post]
// @Success 201 {object} response.Response{} "Successfully Account Created"
// @Failure 400 {object} response.Response{} "Invalid input"
// @Failure 409 {object} response.Response{} "A verified user already exist with given user credentials"
// @Failure 500 {object} response.Response{} "Failed to signup"
func (c *AuthHandler) UserSignUp(ctx *gin.Context) {

	var body request.UserSignUp

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, body)
		return
	}

	var user domain.User
	if err := copier.Copy(&user, body); err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "failed to copy details", err, nil)
		return
	}

	err := c.authUseCase.UserSignUp(ctx, user)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, usecase.ErrUserAlreadyExit) {
			statusCode = http.StatusConflict
		}

		response.ErrorResponse(ctx, statusCode, "Failed to signup", err, nil)
		return
	}

	response.SuccessResponse(ctx, http.StatusCreated, "Successfully Account Created", nil)
}

// UserRenewAccessToken godoc
// @Summary Renew Access Token (User)
// @Description API for user to renew access token using refresh token
// @Security ApiKeyAuth
// @Id UserRenewAccessToken
// @Tags User Authentication
// @Param input body request.RefreshToken{} true "Refresh token"
// @Router /api/auth/renew-access-token [post]
// @Success 200 {object} response.Response{data=response.TokenResponse{}} "Successfully generated access token using refresh token"
// @Failure 400 {object} response.Response{} "Invalid input"
// @Failure 401 {object} response.Response{} "Invalid refresh token"
// @Failure 404 {object} response.Response{} "No session found for the given refresh token"
// @Failure 410 {object} response.Response{} "Refresh token expired"
// @Failure 403 {object} response.Response{} "Refresh token blocked"
// @Failure 500 {object} response.Response{} "Failed generate access token"
func (c *AuthHandler) UserRenewAccessToken() gin.HandlerFunc {
	return c.renewAccessToken(token.User)
}

func (c *AuthHandler) renewAccessToken(tokenUser token.UserType) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var body request.RefreshToken

		if err := ctx.ShouldBindJSON(&body); err != nil {
			response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, body)
			return
		}

		refreshSession, err := c.authUseCase.VerifyAndGetRefreshTokenSession(ctx, body.RefreshToken, tokenUser)

		if err != nil {
			var statusCode int

			switch {
			case errors.Is(err, usecase.ErrInvalidRefreshToken):
				statusCode = http.StatusUnauthorized
			case errors.Is(err, usecase.ErrRefreshSessionNotExist):
				statusCode = http.StatusNotFound
			case errors.Is(err, usecase.ErrRefreshSessionExpired):
				statusCode = http.StatusGone
			case errors.Is(err, usecase.ErrRefreshSessionBlocked):
				statusCode = http.StatusForbidden
			default:
				statusCode = http.StatusInternalServerError
			}
			response.ErrorResponse(ctx, statusCode, "Failed verify refresh token", err, nil)
			return
		}

		accessTokenParams := usecaseInterface.GenerateTokenParams{
			UserID:   refreshSession.UserID,
			UserType: tokenUser,
		}

		accessToken, err := c.authUseCase.GenerateAccessToken(ctx, accessTokenParams)

		if err != nil {
			response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed generate access token", err, nil)
			return
		}
		cookieName := "auth-" + string(tokenUser)
		ctx.SetCookie(cookieName, accessToken, 15*60, "", "", false, true)

		accessTokenRes := response.TokenResponse{
			AccessToken: accessToken,
		}
		response.SuccessResponse(ctx, http.StatusOK, "Successfully generated access token using refresh token", accessTokenRes)
	}
}
