package action

import (
	"github.com/chhongzh/chz_Base_Backend/internal/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service struct {
	waitingActions *utils.Queue[*ActionRequest]
	logger         *zap.Logger
	db             *gorm.DB
}

func New(logger *zap.Logger, db *gorm.DB) *Service {
	return &Service{
		logger:         logger,
		db:             db,
		waitingActions: utils.NewQueue[*ActionRequest](),
	}
}
