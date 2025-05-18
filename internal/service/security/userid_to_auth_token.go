package security

import (
	"time"

	"github.com/chhongzh/chz_Base_Backend/internal/model"
	"github.com/golang-jwt/jwt/v5"
)

func (s *Service) UserIDToAuthToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.AuthToken{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 12)),
		},
		UserID: userID,
	})
	return token.SignedString(s.secret)
}
