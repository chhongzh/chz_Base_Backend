package model

import "github.com/golang-jwt/jwt/v5"

type AccessToken struct {
	jwt.RegisteredClaims

	UserID        string
	ApplicationID string
}
