package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/htan06/echo-messenger-rest-api/internal/apperr"
	"github.com/htan06/echo-messenger-rest-api/internal/module/auth/secure"
)

type AuthenticationService struct {
	userRepo       UserReposiotry
	cacheRepo      CacheRepository
	emailOTPSender EmailOTPSender
	jwtProvider    *secure.JWTProvier
	secureRand     *secure.OTPProvider
}

func NewAuthenticationService(
	userRepo UserReposiotry,
	cacheRepo CacheRepository,
	emailOTPSender EmailOTPSender,
	jwtProvider *secure.JWTProvier,
	secureRand *secure.OTPProvider,
) *AuthenticationService {
	return &AuthenticationService{
		userRepo:       userRepo,
		cacheRepo:      cacheRepo,
		emailOTPSender: emailOTPSender,
		jwtProvider:    jwtProvider,
		secureRand:     secureRand,
	}
}

func (as *AuthenticationService) SendOTP(ctx context.Context, email string) error {
	otp := as.secureRand.RandOTP()

	key := "auth-otp-" + email
	if err := as.cacheRepo.SetIfNotExists(ctx, key, otp, time.Minute*5); err != nil {
		return fmt.Errorf("AuthenticationService[SendOTP]: %w", err)
	}

	go as.emailOTPSender.Send(ctx, email, otp)

	return nil
}

func (as *AuthenticationService) VerifyOTP(ctx context.Context, email string, receivedOtp string) (string, error) {
	key := "auth-otp-" + email
	otp, err := as.cacheRepo.Get(ctx, key)
	if err != nil || otp != receivedOtp {
		return "", apperr.NewAppError(apperr.OTPInvalid)
	}

	user, err := as.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", fmt.Errorf("AuthenticationService[VerifyOTP]: %w", err)
	}

	return as.jwtProvider.GenerateAccessToken(user)
}