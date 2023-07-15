package usecase

import (
	"context"

	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler/response"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/domain"
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

func (c *chatUseCase) SaveMessage(ctx context.Context, message domain.Message) (uint, error) {

	var (
		errChan    = make(chan error, 2)
		receiverID uint
	)

	go func() {
		err := c.chatRepo.SaveMessage(ctx, message)
		errChan <- err
	}()

	go func() {
		rID, err := c.chatRepo.FindReceiverOfChatBySenderID(ctx, message.ChatID, message.SenderID)
		receiverID = rID
		errChan <- err
	}()

	for i := 1; i <= 2; i++ {

		select {
		case err := <-errChan:
			if err != nil {
				return 0, err
			}
		case <-ctx.Done():
			return 0, nil
		}
	}

	return receiverID, nil
}
