package application

import (
	"github.com/chhongzh/chz_Base_Backend/internal/model"
)

func (s *Service) CreateApplication(
	name string,
	desc string,
) (*model.Application, error) {
	application := &model.Application{
		Name: name,
		Desc: desc,
	}

	return s.createFromInstance(application)
}

func (s *Service) CreateApplicationWithApplicationID(
	name string,
	desc string,
	applicationID string,
) (*model.Application, error) {
	application := &model.Application{
		Name:          name,
		Desc:          desc,
		ApplicationID: applicationID,
	}

	return s.createFromInstance(application)
}

func (s *Service) createFromInstance(application *model.Application) (*model.Application, error) {
	if err := s.db.Create(application).Error; err != nil {
		return nil, err
	}
	return application, nil
}
