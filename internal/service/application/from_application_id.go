package application

import "github.com/chhongzh/chz_Base_Backend/internal/model"

func (s *Service) FromApplicationID(applicationID string) (*model.Application, error) {
	// 读取
	var application model.Application
	if err := s.db.First(&application, "application_id = ?", applicationID).Error; err != nil {
		return nil, err
	}
	return &application, nil
}
