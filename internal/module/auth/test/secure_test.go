package test

import (
	"testing"

	"github.com/htan06/echo-messenger-rest-api/internal/module/auth/secure"
)

func TestRandOTP(t *testing.T) {
	otpProvider := secure.NewOTPProvider()

	otp := otpProvider.RandOTP()

	if len(otp) != 6 {
		t.Fail()
	}
}