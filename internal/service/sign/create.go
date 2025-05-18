package sign

import (
	"github.com/chhongzh/chz_Base_Backend/internal/problem"
)

func (s *Service) CreateLoginForApplicationID(applicationID string) (string, error) {
	// 检查是否有剩余
	currentCount := len(s.loginSessionPool)
	if currentCount >= s.maxSignSessionCount {
		return "", problem.ErrSignServiceIsBusy
	}

	// 创建一个 LoginSession
	session := NewLoginSession(applicationID)

	// 插入
	s.loginSessionPool[session.SignSessionID] = session
	s.loginSessionWaitingCountPool[session.SignSessionID] = 0

	return session.SignSessionID, nil
}
