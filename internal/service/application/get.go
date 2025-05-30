package application

import "github.com/chhongzh/chz_Base_Backend/internal/model"

func (s *Service) GetApplicationByApplicationID(applicationID string) (*model.Application, error) {
	var application model.Application
	if err := s.db.First(&application, "application_id = ?", applicationID).Error; err != nil {
		return nil, err
	}
	return &application, nil
}

func (s *Service) GetApplicationByApplicationSecret(applicationSecret string) (*model.Application, error) {
	var application model.Application
	if err := s.db.First(&application, "application_secret = ?", applicationSecret).Error; err != nil {
		return nil, err
	}
	return &application, nil
}
