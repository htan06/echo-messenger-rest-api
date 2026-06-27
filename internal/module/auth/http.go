package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationHandler struct {
	authService *AuthenticationService
}

func NewAuthenticationHandler(authService *AuthenticationService) *AuthenticationHandler {
	return &AuthenticationHandler{
		authService: authService,
	}
}

func (ah *AuthenticationHandler) handlerSendOTP(c *gin.Context) {
	ctx := c.Request.Context()

	var req SendOTPReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
	}

	if err := ah.authService.SendOTP(ctx, req.Email); err != nil {
		fmt.Println("ERROR: ", err.Error())
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}

func (ah *AuthenticationHandler) handlerVerifyOTP(c *gin.Context) {
	ctx := c.Request.Context()

	var req VerifyOTPReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
	}

	token, err := ah.authService.VerifyOTP(ctx, req.Email, req.OTP)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token})
}