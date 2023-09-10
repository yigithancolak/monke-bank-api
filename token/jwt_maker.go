package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTMaker struct {
	secretKey string
}

const minSecretKeySize = 32

func newJWTMaker(secretKey string) (*JWTMaker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size. must be at leat %d characters", minSecretKeySize)
	}

	return &JWTMaker{secretKey}, nil
}

func (maker JWTMaker) CreateToken(email string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(email, duration)

	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, payload, err
}
