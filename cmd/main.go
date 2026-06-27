package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/htan06/echo-messenger-rest-api/internal/config"
	"github.com/htan06/echo-messenger-rest-api/internal/module/auth"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	redisConn := config.GetRedisConn()
	postgresConn := config.GetPostgresConn()
	gmailDialer := config.GetGmailDialer()
	mailAddress := config.GetMailAddress()
	jwtConfig := config.GetJWTConfig()

	router := gin.Default()
	authModule := auth.InitAuthModule(postgresConn, redisConn, gmailDialer, jwtConfig, mailAddress)
	authModule.RegisterRouter(router)

	router.Run()
}