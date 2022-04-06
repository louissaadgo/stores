package token

import (
	"fmt"
	"os"
	"time"

	"github.com/louissaadgo/quiqr/auth/src/models"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

func GeneratePasetoToken(payload interface{}) (string, error) {

	key := os.Getenv("PASETO_KEY")

	if len(key) != chacha20poly1305.KeySize {
		return "", fmt.Errorf("key error")
	}

	paseto := paseto.NewV2()

	return paseto.Encrypt([]byte(key), payload, nil)

}

func VerifyPasetoToken(token string) (models.PasetoTokenPayload, bool) {

	payload := models.PasetoTokenPayload{}

	key := os.Getenv("PASETO_KEY")

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
