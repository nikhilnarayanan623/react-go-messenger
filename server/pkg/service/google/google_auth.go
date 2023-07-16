package google

import (
	"context"
	"errors"
	"fmt"

	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/config"
	"google.golang.org/api/idtoken"
)

var (
	ErrEmailNotVerified = errors.New("email not verified from google")
	ErrTokenExpired     = errors.New("google token expired")
)

type GoogleAuth interface {
	Verify(ctx context.Context, token string) (GoogleUser, error)
}

type googleAuth struct {
	googleClientID string
}

type GoogleUser struct {
	FirstName string
	LastName  string
	Email     string
	Picture   string
	// 	EmailVerified string
}

func NewGoogleAuth(cfg config.Config) GoogleAuth {

	return &googleAuth{
		googleClientID: cfg.GoAuthClientId,
	}
}

func (c *googleAuth) Verify(ctx context.Context, token string) (GoogleUser, error) {

	validator, err := idtoken.NewValidator(ctx)
	if err != nil {
		return GoogleUser{}, err
	}

	payload, err := validator.Validate(ctx, token, c.googleClientID)

	if err != nil {
		return GoogleUser{}, err
	}

	return getUserDetails(payload), nil
}

func getUserDetails(payload *idtoken.Payload) GoogleUser {
	var user GoogleUser

	if firstName, ok := payload.Claims["given_name"].(string); ok {
		fmt.Println(firstName)
		user.FirstName = firstName
	}
	if lastName, ok := payload.Claims["family_name"].(string); ok {
		fmt.Println(lastName)
		user.LastName = lastName
	}
	if email, ok := payload.Claims["email"].(string); ok {
		fmt.Println(email)
		user.Email = email
	}
	if picture, ok := payload.Claims["picture"].(string); ok {
		fmt.Println(picture)
		user.Picture = picture
	}
	return user
}
