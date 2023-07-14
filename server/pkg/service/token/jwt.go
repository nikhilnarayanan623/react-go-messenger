package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/config"
	"github.com/nikhilnarayanan623/server/react-go-messenger/pkg/utils"
)

type jwtAuth struct {
	key string
}

// New TokenAuth
func NewTokenService(cfg config.Config) TokenService {

	return &jwtAuth{
		key: cfg.JwtKey,
	}
}

var (
	ErrInvalidUserType    = errors.New("invalid user type")
	ErrInvalidToken       = errors.New("invalid token")
	ErrFailedToParseToken = errors.New("failed to parse token to claims")
	ErrExpiredToken       = errors.New("token expired")
)

type jwtClaims struct {
	TokenID   string
	UserID    uint
	ExpiresAt time.Time
	Role      UserType
	// jwt.RegisteredClaims
}

// Generate a new JWT token string from token request
func (c *jwtAuth) GenerateToken(req GenerateTokenRequest) (GenerateTokenResponse, error) {

	if req.UsedFor != Admin && req.UsedFor != User {

		return GenerateTokenResponse{}, ErrInvalidUserType
	}

	tokenID := utils.GenerateUniqueString()
	claims := &jwtClaims{
		TokenID:   tokenID,
		UserID:    req.UserID,
		Role:      req.UsedFor,
		ExpiresAt: req.ExpireAt,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var (
		tokenString string
		err         error
	)
	// sign the token by user type
	tokenString, err = token.SignedString([]byte(c.key))

	if err != nil {
		return GenerateTokenResponse{}, fmt.Errorf("failed to sign the token \nerror:%w", err)
	}

	response := GenerateTokenResponse{
		TokenID:     tokenID,
		TokenString: tokenString,
	}

	return response, nil
}

// Verify JWT token string and return TokenResponse
func (c *jwtAuth) VerifyToken(req VerifyTokenRequest) (VerifyTokenResponse, error) {

	if req.UsedFor != Admin && req.UsedFor != User {
		return VerifyTokenResponse{}, ErrInvalidUserType
	}

	token, err := jwt.ParseWithClaims(req.TokenString, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}

		return []byte(c.key), nil
	})

	if err != nil {
		if errors.Is(err, ErrExpiredToken) {
			return VerifyTokenResponse{}, ErrExpiredToken
		}
		return VerifyTokenResponse{}, ErrInvalidToken
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok {
		return VerifyTokenResponse{}, ErrFailedToParseToken
	}

	response := VerifyTokenResponse{
		TokenID: claims.TokenID,
		UserID:  claims.UserID,
	}
	return response, nil
}

// Validate claims
func (c *jwtClaims) Valid() error {
	if time.Since(c.ExpiresAt) > 0 {
		return ErrExpiredToken
	}
	return nil
}
