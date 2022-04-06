package controllers

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	cost = 7
)

func HashPassword(password string) (string, bool) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	if err != nil {
		return "", false
	}

	return string(hashedPassword), true
}

func ValidatePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return false
	}

	return true
}
