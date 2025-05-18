package application

import "github.com/chhongzh/chz_Base_Backend/internal/model"

func (s *Service) DeleteApplication(applicationID string) error {
	return s.db.Delete(&model.Application{}, "application_id = ?", applicationID).Error
}
