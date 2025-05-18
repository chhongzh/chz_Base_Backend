package security

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	logger *zap.Logger
	db     *gorm.DB

	secret []byte
}

func New(logger *zap.Logger, db *gorm.DB, secret []byte) *Service {
	return &Service{
		logger: logger,
		db:     db,

		secret: secret,
	}
}
