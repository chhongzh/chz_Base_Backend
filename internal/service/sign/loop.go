package sign

import (
	"time"

	"github.com/chhongzh/chz_Base_Backend/internal/problem"
	"go.uber.org/zap"
)

func (s *Service) Loop() {
	s.logger.Info("Sign Service 线程")

	// 每半分钟检查是否过期
	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for range ticker.C {
		s.cleanUpLoop()
	}
}

func (s *Service) cleanUpLoop() {
	cleaned := 0

	for sessionID, session := range s.loginSessionPool {
		if time.Since(session.CreatedAt) > time.Minute*1 { // 1分钟没有处理就删除
			// 先进行释放
			session.Emit("", problem.ErrSignSessionExpired)

			s.logger.Info("清理登录会话",
				zap.String("Session ID", sessionID),
			)

			delete(s.loginSessionPool, sessionID)

			cleaned++
		}
	}

	if cleaned > 0 {
		s.logger.Info("本次清理登录会话",
			zap.Int("清理数量", cleaned),
		)
	}
}
