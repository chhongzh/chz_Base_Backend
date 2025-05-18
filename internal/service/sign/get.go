package sign

import (
	"github.com/chhongzh/chz_Base_Backend/internal/problem"
)

func (s *Service) GetSignSession(signSessionID string) (*SignSession, error) {
	session, ok := s.loginSessionPool[signSessionID]
	if !ok {
		return nil, problem.ErrSignSessionNotFound
	}
	return session, nil
}

func (s *Service) GetSignSessionWaitingCount(signSessionID string) (int, error) {
	session, ok := s.loginSessionWaitingCountPool[signSessionID]
	if !ok {
		return 0, problem.ErrSignSessionNotFound
	}
	return session, nil
}
