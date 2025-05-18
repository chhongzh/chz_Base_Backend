package security

import (
	"time"

	"github.com/chhongzh/chz_Base_Backend/internal/model"
	"github.com/golang-jwt/jwt/v5"
)

func (s *Service) ApplicationIDAndUserIDToAccessToken(applicationID, userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.AccessToken{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 12)),
		},
		ApplicationID: applicationID,
		UserID:        userID,
	})
	return token.SignedString(s.secret)
}
