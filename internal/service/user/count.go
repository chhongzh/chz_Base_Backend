package user

import "github.com/chhongzh/chz_Base_Backend/internal/model"

func (s *Service) Count() (int64, error) {
	var count int64
	return count, s.db.Model(&model.User{}).Count(&count).Error
}
