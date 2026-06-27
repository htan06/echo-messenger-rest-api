package secure

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/htan06/echo-messenger-rest-api/internal/config"
	"github.com/htan06/echo-messenger-rest-api/internal/module/auth/model"
)

type UserClaimsAccess struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	jwt.RegisteredClaims
}

type JWTProvier struct {
	cfg *config.JWTConfig
}

func NewJWTProvider(cfg *config.JWTConfig) *JWTProvier {
	return &JWTProvier{
		cfg: cfg,
	}
}

func (jp *JWTProvier) GenerateAccessToken(user model.User) (string, error) {
	claim := UserClaimsAccess{
		FirstName: user.FirstName,
		LastName:  *user.LastName,
		Username:  user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.Email,
			Issuer:    "echo-authenticator",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jp.cfg.TtlAccess())),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	tokenSigned, err := token.SignedString(jp.cfg.PrivateKeyAccess())

	if err != nil {
		return "", fmt.Errorf("JWTProvider[GenerateAccessToken]: %w", err)
	}

	return tokenSigned, nil
}
