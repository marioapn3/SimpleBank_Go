package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrExpiredToken = errors.New("token has invalid claims: token is expired")
	ErrInvalidToken = errors.New("token is invalid")
)

// Payload adalah struktur data yang akan disematkan ke dalam token JWT
type Payload struct {
	ID                   uuid.UUID `json:"id"`
	Username             string    `json:"username"`
	jwt.RegisteredClaims           // Embed jwt.RegisteredClaims to satisfy jwt.Claims interface
}

// NewPayload membuat payload baru
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:       tokenID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiresAt.Time) {
		return ErrExpiredToken
	}
	return nil
}
