package jwtmanager

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessClaims struct {
	jwt.RegisteredClaims
	BirthDate time.Time `json:"birth_date"`
}
