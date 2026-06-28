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

func (ah *AuthenticationHandler) handleSendOTP(c *gin.Context) {
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

func (ah *AuthenticationHandler) handleVerifyOTP(c *gin.Context) {
	ctx := c.Request.Context()

	var req VerifyOTPReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
	}

	tokenResp, err := ah.authService.VerifyOTP(ctx, req.Email, req.OTP)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, tokenResp)
}

func (ah *AuthenticationHandler) handleRegisterUser(c *gin.Context) {
	ctx := c.Request.Context()

	var req RegisterUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Status(http.StatusBadRequest)
	}

	tokenResp, err := ah.authService.RegisterUser(ctx, req)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, tokenResp)
}