package application

import (
	"github.com/chhongzh/chz_Base_Backend/internal/model"
	"github.com/chhongzh/chz_Base_Backend/internal/utils"
)

func (s *Service) ListApplication(page int) ([]*model.Application, error) {
	var res []*model.Application
	err := utils.Pagination(s.db.Order("created_at DESC"), page, 15).Find(&res).Error
	return res, err
}
