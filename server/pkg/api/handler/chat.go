package handler

import (
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/api/handler/interfaces"
	usecase "github.com/nikhilnarayanan623/server/react-go-messenger/pkg/usecase/interfaces"
)

type chatHandler struct {
	usecase usecase.ChatUseCase
}

func NewChatHandler(usecase usecase.ChatUseCase) interfaces.ChatHandler {
	return &chatHandler{
		usecase: usecase,
	}
}
