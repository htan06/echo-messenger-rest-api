package auth

type SendOTPReq struct {
	Email string `json:"email"`
}

type VerifyOTPReq struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

type TokenResp struct {
	ExistsUser bool `json:"exists_user"`
	Tokens     map[string]string
}

type RegisterUserReq struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Username      string `json:"username"`
	PhoneNumber   string `json:"phone_number"`
	RegisterToken string `json:"register_token"`
}
