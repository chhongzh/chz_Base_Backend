package action

import (
	"github.com/chhongzh/chz_Base_Backend/internal/model"
	"go.uber.org/zap"
)

func (s *Service) Loop() {
	s.logger.Info("Action Service 线程")

	for {
		req := s.waitingActions.Dequeue()

		// 插入数据库
		err := s.db.Create(&model.Action{
			FromApplicationID: req.FromApplicationID,
			Message:           req.Message,
			Meta:              req.Meta,
		}).Error
		if err != nil { // 错误提示
			s.logger.Warn(
				"Action 插入数据库失败",
				zap.Any("Request", req),
				zap.Error(err),
			)
		}
	}
}
