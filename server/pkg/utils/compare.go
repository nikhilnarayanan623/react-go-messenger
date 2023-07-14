package utils

import (
	"errors"

	"github.com/nikhilnarayanan623/go-socket-chat/server/pkg/domain"
)

// To compare and return error for the same fields
func CompareUserExistingDetails(user1, user2 domain.User) error {
	var err error
	if user1.Email == user2.Email {
		err = AppendMessageToError(err, "user already exist with this email")
	}
	if user1.UserName == user2.UserName {
		err = AppendMessageToError(err, "user already exist with this user name")
	}

	if err == nil {
		return errors.New("failed to find existing details")
	}

	return err
}
