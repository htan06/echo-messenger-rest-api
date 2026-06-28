package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/htan06/echo-messenger-rest-api/internal/config"
	"github.com/htan06/echo-messenger-rest-api/internal/module/auth"
	"github.com/joho/godotenv"
)

func main() {
	wd, _ := os.Getwd()
	fmt.Println("cwd:", wd)
	
	err := godotenv.Load()
	if err != nil {
		log.Println("WAR: Cannot load .env file")
	}

	privateKeyPath := os.Getenv("PRIVATE_KEY_PATH")
	privateData, err := os.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatal("Cannot load jwt private key",  err.Error())
	}

	publicKeyPath := os.Getenv("PUBLIC_KEY_PATH")
	publicData, err := os.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatal("Cannot load jwt public key", err.Error())
	}
	
	jwtConfig := config.GetJWTConfig(publicData, privateData)

	redisConn := config.GetRedisConn()
	postgresConn := config.GetPostgresConn()
	gmailDialer := config.GetGmailDialer()
	mailAddress := config.GetMailAddress()

	router := gin.Default()
	authModule := auth.InitAuthModule(postgresConn, redisConn, gmailDialer, jwtConfig, mailAddress)
	authModule.RegisterRouter(router)

	router.Run()
}