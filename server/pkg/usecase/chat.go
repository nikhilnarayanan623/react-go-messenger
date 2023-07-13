package usecase

import (
	repoInterface "github.com/nikhilnarayanan623/go-socket-chat/server/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/go-socket-chat/server/pkg/usecase/interfaces"
)

type chatUseCase struct {
	fileRepo repoInterface.ChatRepository
}

func NewChatUseCase(fileRepo repoInterface.ChatRepository) interfaces.ChatUseCase {
	return &chatUseCase{
		fileRepo: fileRepo,
	}
}
