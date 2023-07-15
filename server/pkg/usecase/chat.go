package usecase

import (
	"context"

	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler/response"
	repoInterface "github.com/nikhilnarayanan623/react-go-messenger/server/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/utils"
)

type chatUseCase struct {
	chatRepo repoInterface.ChatRepository
}

func NewChatUseCase(chatRepo repoInterface.ChatRepository) interfaces.ChatUseCase {
	return &chatUseCase{
		chatRepo: chatRepo,
	}
}

func (c *chatUseCase) FindAllRecentChatsOfUser(ctx context.Context, userID uint, pagination request.Pagination) ([]response.Chat, error) {

	chats, err := c.chatRepo.FindAllRecentChatsOfUser(ctx, userID, pagination)
	if err != nil {
		return nil, utils.PrependMessageToError(err, "failed to get user chats from database")
	}

	return chats, nil
}

func (c *chatUseCase) SaveChat(ctx context.Context, user1ID, user2ID uint) (uint, error) {

	// check users have already a chat exist
	chatID, err := c.chatRepo.FindChatIDByUser1AndUser2ID(ctx, user1ID, user2ID)
	if err != nil {
		return 0, utils.PrependMessageToError(err, "failed to check chat is already in database")
	}

	if chatID != 0 {
		return chatID, nil
	}

	// if users have no chat then create it
	chatID, err = c.chatRepo.SaveChat(ctx, user1ID, user2ID)
	if err != nil {
		return 0, utils.PrependMessageToError(err, "failed to save chat on database")
	}

	return chatID, nil
}

func (c *chatUseCase) FindAllMessagesOfUserForAChat(ctx context.Context,
	chatID, userID uint, pagination request.Pagination) ([]response.Message, error) {

	messages, err := c.chatRepo.FindAllMessagesByChatAndUserID(ctx, chatID, userID, pagination)
	if err != nil {
		return nil, utils.PrependMessageToError(err, "failed to get message of the chat from database")
	}

	return messages, nil
}
