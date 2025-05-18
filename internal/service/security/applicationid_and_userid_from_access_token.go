package security

import (
	"github.com/chhongzh/chz_Base_Backend/internal/model"
	"github.com/chhongzh/chz_Base_Backend/internal/problem"
	"github.com/golang-jwt/jwt/v5"
)

func (s *Service) ApplicationIDAndUserIDFromAccessToken(token string) (string, string, error) {
	claims, err := jwt.ParseWithClaims(token, &model.AccessToken{}, func(t *jwt.Token) (interface{}, error) {
		return s.secret, nil
	})
	if err != nil {
		return "", "", err
	}

	if !claims.Valid {
		return "", "", problem.ErrInvalidAccessToken
	}

	accessToken, ok := claims.Claims.(*model.AccessToken)
	if !ok {
		return "", "", problem.ErrInvalidAccessToken
	}

	return accessToken.ApplicationID, accessToken.UserID, nil
}
