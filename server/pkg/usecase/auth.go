package usecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/api/handler/request"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/domain"
	repo "github.com/nikhilnarayanan623/react-go-messenger/server/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/service/google"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/service/token"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/react-go-messenger/server/pkg/utils"
)

type authUseCase struct {
	authRepo     repo.AuthRepository
	googleAuth   google.GoogleAuth
	userRepo     repo.UserRepository
	tokenService token.TokenService
}

func NewAuthUseCase(authRepo repo.AuthRepository, tokenService token.TokenService,
	googleAuth google.GoogleAuth, userRepo repo.UserRepository) interfaces.AuthUseCase {

	return &authUseCase{
		authRepo:     authRepo,
		googleAuth:   googleAuth,
		tokenService: tokenService,
		userRepo:     userRepo,
	}
}

const (
	AccessTokenDuration  = time.Minute * 20
	RefreshTokenDuration = time.Hour * 24 * 7
)

func (c *authUseCase) UserLogin(ctx context.Context, loginDetails request.Login) (userID uint, err error) {

	var user domain.User
	if loginDetails.Email != "" {
		user, err = c.userRepo.FindUserByEmail(ctx, loginDetails.Email)
	} else if loginDetails.UserName != "" {
		user, err = c.userRepo.FindUserByUserName(ctx, loginDetails.UserName)
	} else {
		return userID, ErrEmptyLoginCredentials
	}

	if err != nil {
		return userID, fmt.Errorf("failed to find user \nerror: %w", err)
	}

	if user.ID == 0 {
		return userID, ErrUserNotExist
	}

	if !utils.VerifyHashAndPassword(user.Password, loginDetails.Password) {
		return userID, ErrWrongPassword
	}

	return user.ID, nil
}

func (c *authUseCase) GoogleLogin(ctx context.Context, token string) (userID uint, err error) {

	googleUser, err := c.googleAuth.Verify(ctx, token)
	if err != nil {
		return 0, err
	}

	existUser, err := c.userRepo.FindUserByEmail(ctx, googleUser.Email)
	if err != nil {
		return userID, fmt.Errorf("failed to get user details with given email \nerror:%v", err.Error())
	}

	if existUser.ID != 0 {
		return existUser.ID, nil
	}

	// user := domain.User{
	// 	GoogleImage: googleUser.Picture,
	// 	FirstName: ,
	// }

	// // create a random user name for user based on user name
	// user.UserName = utils.GenerateRandomUserName(user.FirstName)

	// userID, err = c.userRepo.SaveUser(ctx, user)
	// if err != nil {
	// 	return userID, fmt.Errorf("failed to save user details \nerror:%v", err.Error())
	// }

	return userID, nil
}

func (c *authUseCase) GenerateAccessToken(ctx context.Context, tokenParams interfaces.GenerateTokenParams) (string, error) {

	tokenReq := token.GenerateTokenRequest{
		UserID:   tokenParams.UserID,
		UsedFor:  tokenParams.UserType,
		ExpireAt: time.Now().Add(AccessTokenDuration),
	}

	tokenRes, err := c.tokenService.GenerateToken(tokenReq)

	return tokenRes.TokenString, err
}
func (c *authUseCase) GenerateRefreshToken(ctx context.Context, tokenParams interfaces.GenerateTokenParams) (string, error) {

	expireAt := time.Now().Add(RefreshTokenDuration)
	tokenReq := token.GenerateTokenRequest{
		UserID:   tokenParams.UserID,
		UsedFor:  tokenParams.UserType,
		ExpireAt: expireAt,
	}
	tokenRes, err := c.tokenService.GenerateToken(tokenReq)
	if err != nil {
		return "", err
	}

	err = c.authRepo.SaveRefreshSession(ctx, domain.RefreshSession{
		UserID:       tokenParams.UserID,
		TokenID:      tokenRes.TokenID,
		RefreshToken: tokenRes.TokenString,
		ExpireAt:     expireAt,
	})
	if err != nil {
		return "", err
	}
	log.Printf("successfully refresh token created and refresh session stored in database")
	return tokenRes.TokenString, nil
}

func (c *authUseCase) VerifyAndGetRefreshTokenSession(ctx context.Context, refreshToken string, usedFor token.UserType) (domain.RefreshSession, error) {

	verifyReq := token.VerifyTokenRequest{
		TokenString: refreshToken,
		UsedFor:     usedFor,
	}
	verifyRes, err := c.tokenService.VerifyToken(verifyReq)
	if err != nil {
		return domain.RefreshSession{}, utils.PrependMessageToError(ErrInvalidRefreshToken, err.Error())
	}

	refreshSession, err := c.authRepo.FindRefreshSessionByTokenID(ctx, verifyRes.TokenID)
	if err != nil {
		return refreshSession, err
	}

	if refreshSession.TokenID == "" {
		return refreshSession, ErrRefreshSessionNotExist
	}

	if time.Since(refreshSession.ExpireAt) > 0 {
		return domain.RefreshSession{}, ErrRefreshSessionExpired
	}

	if refreshSession.IsBlocked {
		return domain.RefreshSession{}, ErrRefreshSessionBlocked
	}

	return refreshSession, nil
}

func (c *authUseCase) UserSignUp(ctx context.Context, signUpDetails domain.User) error {

	existUser, err := c.userRepo.FindUserByUserNameEmailNotID(ctx, signUpDetails)
	if err != nil {
		return utils.PrependMessageToError(err, "failed to check user details already exist")
	}

	// if user already exist then find the exists details and append that error with user already exist error
	if existUser.ID != 0 {
		err = utils.CompareUserExistingDetails(existUser, signUpDetails)
		err = utils.AppendMessageToError(ErrUserAlreadyExit, err.Error())
		return err
	}

	hashPass, err := utils.GenerateHashFromPassword(signUpDetails.Password)
	if err != nil {
		return utils.PrependMessageToError(err, "failed to hash the password")
	}

	signUpDetails.Password = string(hashPass)
	_, err = c.userRepo.SaveUser(ctx, signUpDetails)

	if err != nil {
		return utils.PrependMessageToError(err, "failed to save user details")
	}
	return nil
}
