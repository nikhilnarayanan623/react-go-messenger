package usecase

import (
	repoInterface "github.com/nikhilnarayanan623/server/react-go-messenger/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/usecase/interfaces"
)

type chatUseCase struct {
	chatRepo repoInterface.ChatRepository
}

func NewChatUseCase(chatRepo repoInterface.ChatRepository) interfaces.ChatUseCase {
	return &chatUseCase{
		chatRepo: chatRepo,
	}
}
