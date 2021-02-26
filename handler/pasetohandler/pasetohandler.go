package pasetohandler

import (
	"log"
	"net/http"
	"time"

	"github.com/o1egl/paseto"
)

func PasetoSignIn(w http.ResponseWriter, r *http.Request) {
	symetricKey := []byte("YELLOW SUBMARINE, BLACK WIZARDRY")
	now := time.Now()
	exp := now.Add(time.Minute * 5)
	nbt := now

	jsonToken := paseto.JSONToken{
		Audience:   "private",
		Issuer:     "go-web",
		Jti:        "123",
		Subject:    "test_subject",
		IssuedAt:   now,
		Expiration: exp,
		NotBefore:  nbt,
	}
	jsonToken.Set("data", "Test Data")
	footer := "PT Jojo Nomic Indonesia"

	// Encrypt the data
	// token, err := paseto.Encrypt(symetricKey, jsonToken, footer)
	token, err := paseto.NewV2().Encrypt(symetricKey, jsonToken, footer)
	if err != nil {
		log.Printf("Error token Encryption: %s, key length: %d", err, len(symetricKey))
	}
	log.Print(token)

	http.SetCookie(w, &http.Cookie{
		Name:    "pasetoToken",
		Value:   token,
		Expires: exp,
	})
}
