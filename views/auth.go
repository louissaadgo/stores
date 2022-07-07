package views

type Auth struct {
	AuthToken string `json:"auth_token"`
}

type UserAuthData struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Country string `json:"country"`
	Status  string `json:"status"`
	Image   string `json:"image"`
}

type UserAuth struct {
	AuthToken string       `json:"auth_token"`
	UserData  UserAuthData `json:"user_data"`
}
