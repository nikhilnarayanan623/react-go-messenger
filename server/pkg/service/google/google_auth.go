package google

import (
	"context"
	"errors"

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

	gUser := GoogleUser{
		FirstName: payload.Claims["given_name"].(string),
		LastName:  payload.Claims["family_name"].(string),
		Email:     payload.Claims["email"].(string),
		Picture:   payload.Claims["picture"].(string),
	}

	return gUser, nil
}
