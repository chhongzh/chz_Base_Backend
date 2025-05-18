package user

import "github.com/chhongzh/chz_Base_Backend/internal/model"

func (s *Service) FromUsername(username string) (*model.User, error) {
	var user model.User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Service) FromUserID(userID string) (*model.User, error) {
	var user model.User
	if err := s.db.Where("user_id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
