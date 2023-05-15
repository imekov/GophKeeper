package auth

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	key []byte
}

func NewJWT() JWT {

	secret := make([]byte, 256)
	_, err := rand.Read(secret)
	if err != nil {
		log.Fatal(err)
	}

	return JWT{
		key: secret,
	}
}

func (j JWT) Create(ttl time.Duration, userID int) (string, error) {
	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["dat"] = userID
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.key)
	if err != nil {
		return "", fmt.Errorf("\nошибка подписи токена: %w", err)
	}

	return token, nil
}

func (j JWT) Validate(token string) (int, error) {
	t, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("\nошибка при проверке токена: %s", jwtToken.Header["alg"])
		}

		return j.key, nil
	})
	if err != nil {
		return 0, fmt.Errorf("\nошибка при проверке токена: %w", err)
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return 0, fmt.Errorf("\nтокен не валидный")
	}

	return int(claims["dat"].(float64)), nil
}
