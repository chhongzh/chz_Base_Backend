package announcement

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	logger *zap.Logger
	db     *gorm.DB
}

func NewService(logger *zap.Logger, db *gorm.DB) *Service {
	return &Service{
		logger: logger,
		db:     db,
	}
}
