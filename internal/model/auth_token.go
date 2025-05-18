package model

import "github.com/golang-jwt/jwt/v5"

type AuthToken struct {
	jwt.RegisteredClaims

	UserID string
}
