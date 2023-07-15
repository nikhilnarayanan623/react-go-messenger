package google

import (
	"context"
	"errors"
	"fmt"
	"strings"

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
	fmt.Println("id: ", c.googleClientID)
	payload, err := validator.Validate(ctx, token, c.googleClientID)

	if err != nil {
		return GoogleUser{}, err
	}

	name := payload.Claims["given_name"].(string)
	email := payload.Claims["email"].(string)
	picture := payload.Claims["picture"].(string)

	// separate first name and last name from the name
	firstName, ok1 := strings.CutSuffix(name, " ")
	lastName, ok2 := strings.CutPrefix(name, " ")

	if !ok1 && !ok2 { // failed to cut first name and last then assign first only and last name as empty
		firstName, lastName = name, ""
	}

	user := GoogleUser{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Picture:   picture,
	}
	fmt.Println(user.FirstName, ":", user.LastName, user.Email, user.Picture)
	return user, nil
}
