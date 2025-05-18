package sign

import (
	"go.uber.org/zap"
)

type Service struct {
	loginSessionPool             map[string]*SignSession
	loginSessionWaitingCountPool map[string]int
	logger                       *zap.Logger

	maxSignSessionCount        int
	maxSignSessionWaitingCount int
}

func New(logger *zap.Logger, maxSignSessionCount int, maxSignSessionWaitingCount int) *Service {
	service := &Service{
		loginSessionPool:             make(map[string]*SignSession),
		loginSessionWaitingCountPool: make(map[string]int),
		logger:                       logger,
		maxSignSessionCount:          maxSignSessionCount,
		maxSignSessionWaitingCount:   maxSignSessionWaitingCount,
	}

	return service
}

func (s *Service) GetLoginSessionCount() int {
	return len(s.loginSessionPool)
}
