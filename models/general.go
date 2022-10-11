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

type UserImageAndName struct {
	Image string `json:"image"`
	Name  string `json:"name"`
}
