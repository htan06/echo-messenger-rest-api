package auth

type SendOTPReq struct {
	Email string `json:"email"`
}

type VerifyOTPReq struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}
