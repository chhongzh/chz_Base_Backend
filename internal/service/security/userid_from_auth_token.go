package security

import (
	"github.com/chhongzh/chz_Base_Backend/internal/model"
	"github.com/chhongzh/chz_Base_Backend/internal/problem"
	"github.com/golang-jwt/jwt/v5"
)

func (s *Service) UserIDFromAuthToken(token string) (string, error) {
	claims, err := jwt.ParseWithClaims(token, &model.AuthToken{}, func(t *jwt.Token) (interface{}, error) {
		return s.secret, nil
	})
	if err != nil {
		return "", err
	}

	if !claims.Valid {
		return "", problem.ErrInvalidAuthToken
	}

	authToken, ok := claims.Claims.(*model.AuthToken)
	if !ok {
		return "", problem.ErrInvalidAuthToken
	}

	return authToken.UserID, nil
}
