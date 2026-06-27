package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/htan06/echo-messenger-rest-api/internal/config"
	"github.com/htan06/echo-messenger-rest-api/internal/module/auth/infra"
	"github.com/htan06/echo-messenger-rest-api/internal/module/auth/secure"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"gopkg.in/gomail.v2"
)

type AuthModule struct {
	authHandler *AuthenticationHandler
}

func InitAuthModule(
	postgresConn *pgxpool.Pool,
	redisConn *redis.Client,
	dialer *gomail.Dialer,
	jwtConfig *config.JWTConfig,
	mailAddress *config.MailAddress,
) *AuthModule {

	userRepo := infra.NewPostgresUserRepository(postgresConn)
	cacheRepository := infra.NewRedisCacheRepository(redisConn)
	emailOTPSender := infra.NewGmailOTPSender(dialer, mailAddress)

	otpProvider := secure.NewOTPProvider()
	jwtProvider := secure.NewJWTProvider(jwtConfig)

	authService := NewAuthenticationService(userRepo, cacheRepository, emailOTPSender, jwtProvider, otpProvider)

	authHandler := NewAuthenticationHandler(authService)

	return &AuthModule{
		authHandler: authHandler,
	}
}

func (am *AuthModule) RegisterRouter(r *gin.Engine, middlewares...gin.HandlerFunc) {
	group := r.Group("/auth")

	group.POST("/send-otp", am.authHandler.handlerSendOTP)
	group.POST("/verify-otp", am.authHandler.handlerVerifyOTP)
}
