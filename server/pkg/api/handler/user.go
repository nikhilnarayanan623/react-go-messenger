package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler/interfaces"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler/response"
	usecaseInterface "github.com/nikhilnarayanan623/server/react-go-messenger/pkg/usecase/interfaces"
)

type userHandler struct {
	usecase usecaseInterface.UserUseCase
}

func NewUserHandler(usecase usecaseInterface.UserUseCase) interfaces.UserHandler {
	return &userHandler{
		usecase: usecase,
	}
}

// ListUsers godoc
// @Summary List all users (User)
// @Description API for user to list all users in the application
// @Security ApiKeyAuth
// @Id ListUsers
// @Tags Users
// @Param page_number query int false "Page Number"
// @Param count query int false "Count"
// @Router /api/users/all [get]
// @Success 200 {object} response.Response{data=[]response.User} "Successfully retrieved all users"
// @Success 204 {object} response.Response{} "There is no users"
// @Failure 500 {object} response.Response{} "Failed retrieved all users"
func (c *userHandler) ListUsers(ctx *gin.Context) {

	pagination := request.GetPagination(ctx)

	users, err := c.usecase.FindAllUsers(ctx, pagination)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed retrieved all users", err, nil)
		return
	}

	if len(users) == 0 {
		response.SuccessResponse(ctx, http.StatusNoContent, "There is no users", nil)
		return
	}

	response.SuccessResponse(ctx, http.StatusOK, "Successfully retrieved all users", users)
}
