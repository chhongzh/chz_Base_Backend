package sign

import (
	"github.com/chhongzh/chz_Base_Backend/internal/problem"
	"go.uber.org/zap"
)

func (s *Service) WaitForSignSession(signSessionID string) (string, error) {
	// 检查是否有太多正在等待的
	waitingCount, err := s.GetSignSessionWaitingCount(signSessionID)
	if err != nil {
		return "", err
	}
	if waitingCount >= s.maxSignSessionWaitingCount {
		s.logger.Info("Sign 等待请求 Dropped, 因为此前有太多等待数量", zap.String("Session ID", signSessionID))
		return "", problem.ErrSignSessionTooManyWaiting
	}
	s.loginSessionWaitingCountPool[signSessionID]++

	// 输出
	s.logger.Info("开始监听", zap.String("Session ID", signSessionID), zap.Int("此前监听数量", waitingCount))

	session, err := s.GetSignSession(signSessionID)
	if err != nil {
		return "", err
	}

	return session.Wait()
}
