package auth

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/rs/zerolog"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	key []byte
}

// NewJWT создает новый секретный ключ для работы JWT.
func NewJWT(logger *zerolog.Logger) JWT {

	secret := make([]byte, 256)
	_, err := rand.Read(secret)
	if err != nil {
		newPrint := fmt.Sprintf("unable to create new JWT secret : %s", err.Error())
		logger.Error().Msg(newPrint)
	}

	return JWT{
		key: secret,
	}
}

// Create записывает пользовательский идентификатор в токен и возвращает его.
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

// Validate проверяет корректность токена.
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
