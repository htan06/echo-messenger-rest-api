package test

import (
	"testing"

	"github.com/htan06/echo-messenger-rest-api/internal/security"
)

func TestRandOTP(t *testing.T) {
	otpProvider := security.NewOTPProvider()

	otp := otpProvider.RandOTP()

	if len(otp) != 6 {
		t.Fail()
	}
}