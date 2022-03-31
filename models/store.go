package models

type Store struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	CountryCode string `json:"country_code"`
	Phone       string `json:"phone"`
	PublicEmail string `json:"public_email"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type StoreSignup struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	CountryCode string `json:"country_code"`
	Phone       string `json:"phone"`
	PublicEmail string `json:"public_email"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type StoreResponse struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	CountryCode string `json:"country_code"`
	Phone       string `json:"phone"`
	PublicEmail string `json:"public_email"`
	Email       string `json:"email"`
}
