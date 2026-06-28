package user

import (
	"github.com/gin-gonic/gin"
	"github.com/htan06/echo-messenger-rest-api/internal/module/user/infra"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserModule struct {
	userHandler *UserHandler
}

func InitUserModule(
	postgresConn *pgxpool.Pool,
) *UserModule {
	userRepo := infra.NewPostgresUserRepository(postgresConn)
	
	userService := NewUserService(userRepo)

	userHandler := NewUserHandler(userService)

	return &UserModule{
		userHandler: userHandler,
	}
}

func (um *UserModule) RegisterRouter(r *gin.RouterGroup, requireAccessTokenMiddleware gin.HandlerFunc) {
	user := r.Group("/user")

	user.GET("/me", requireAccessTokenMiddleware, um.userHandler.HandleGetInfo)
	user.PATCH("/me/info", requireAccessTokenMiddleware, um.userHandler.HandleUpdateInfo)
	user.PATCH("/me/read-status", requireAccessTokenMiddleware, um.userHandler.HandleChangeReadStatus)
	user.PATCH("/me/username", requireAccessTokenMiddleware, um.userHandler.HandleUpdateUsername)
}