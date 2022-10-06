package models

type OTP struct {
	OTPToken string `json:"otp"`
}

type PhoneNumber struct {
	Phone string `json:"phone"`
}

type OTPAndPhoneAndPassword struct {
	OTPToken string `json:"otp"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
