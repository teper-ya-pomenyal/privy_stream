package jwtmanager

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Manager struct {
	privateKey *rsa.PrivateKey
	accessTTL  time.Duration
}

func New(privateKey *rsa.PrivateKey, accessTTL time.Duration) *Manager {
	return &Manager{privateKey: privateKey, accessTTL: accessTTL}
}

func (m *Manager) NewAccessToken(userUUID string) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   userUUID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.accessTTL)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(m.privateKey)
}

func (m *Manager) NewRefreshToken() (string, error) {
	randBytes := make([]byte, 32)
	if _, err := rand.Read(randBytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(randBytes), nil
}
