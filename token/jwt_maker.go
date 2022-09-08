package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const minSecretKey = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (*JWTMaker, error) {
	if len(secretKey) < minSecretKey {
		return nil, errors.New("the secret key is too short.")
	}

	return &JWTMaker{
		secretKey,
	}, nil
}

// CreateToken create a new token for a specific username and duration.
func (maker *JWTMaker) CreateToken(username string, ttl time.Duration) (string, error) {
	payload, err := NewPayload(username, ttl)
	if err != nil {
		return "", err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return jwtToken.SignedString([]byte(maker.secretKey))
}

// VerifyToken Checks if the provided token is valid.
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {

	// to check if the token has the a same encryption algorithm as mine.

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // since we are using SigningMethodHS256 that is instance of SigningMethodHMAC
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)

	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiryToken) {
			return nil, ErrExpiryToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)

	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil

}