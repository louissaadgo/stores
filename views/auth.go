package views

type UserAuthData struct {
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	Country       string `json:"country"`
	Status        string `json:"status"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	VerifiedPhone bool   `json:"verified_phone"`
}

type UserAuth struct {
	AuthToken string       `json:"auth_token"`
	UserData  UserAuthData `json:"user_data"`
}

type AuthWeb struct {
	AuthToken string `json:"auth_token"`
}

type CurrentTypeWeb struct {
	CurrentType string `json:"current_type"`
	Name        string `json:"name"`
}
