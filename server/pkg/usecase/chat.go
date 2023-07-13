package usecase

import (
	repoInterface "github.com/nikhilnarayanan623/go-socket-chat/server/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/go-socket-chat/server/pkg/usecase/interfaces"
)

type chatUseCase struct {
	chatRepo repoInterface.ChatRepository
}

func NewChatUseCase(chatRepo repoInterface.ChatRepository) interfaces.ChatUseCase {
	return &chatUseCase{
		chatRepo: chatRepo,
	}
}
