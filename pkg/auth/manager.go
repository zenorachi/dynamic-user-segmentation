package auth

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/form3tech-oss/jwt-go"
)

type TokenManager interface {
	NewJWT(userId int, ttl time.Duration) (string, error)
	NewRefreshToken() (string, error)
	ParseToken(accessToken string) (int, error)
}

type Manger struct {
	secret string
}

func NewManager(secret string) *Manger {
	return &Manger{secret: secret}
}

func (m *Manger) NewJWT(userId int, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(ttl).Unix(),
		Subject:   strconv.Itoa(userId),
	})

	return token.SignedString([]byte(m.secret))
}

func (m *Manger) NewRefreshToken() (string, error) {
	buff := make([]byte, 32)

	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)

	if _, err := r.Read(buff); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", buff), nil
}

func (m *Manger) ParseToken(accessToken string) (int, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.secret), nil
	})
	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("unable to get user claims from token")
	}

	id, err := strconv.Atoi(claims["sub"].(string))
	if err != nil {
		return 0, errors.New("invalid subject")
	}

	return id, nil
}
