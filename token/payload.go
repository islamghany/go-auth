package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrExpiryToken  = errors.New("token has expired")
	ErrInvalidToken = errors.New("invalid token")
)

// Payload contains the payload data for the token.
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, ttl time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}
	return &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(ttl),
	}, nil
}

func (p *Payload) Valid() error {

	if time.Now().After(p.ExpiredAt) {
		return ErrExpiryToken
	}

	return nil
}
