package token

import (
	"fmt"
	"stores/models"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

const (
	key = "01234567890123456789012345678924"
)

func GeneratePasetoToken(id, userID, userType string) (string, error) {

	if len(key) != chacha20poly1305.KeySize {
		return "", fmt.Errorf("key error")
	}

	payload := models.PasetoTokenPayload{
		ID:       id,
		UserID:   userID,
		UserType: userType,
		IssuedAt: time.Now(),
	}

	paseto := paseto.NewV2()

	return paseto.Encrypt([]byte(key), payload, nil)

}

func VerifyPasetoToken(token string) (models.PasetoTokenPayload, bool) {

	payload := models.PasetoTokenPayload{}

	paseto := paseto.NewV2()

	err := paseto.Decrypt(token, []byte(key), &payload, nil)
	if err != nil {
		return payload, false
	}

	if !time.Now().Before(payload.ExpiresAt) {
		return payload, false
	}

	return payload, true
}
