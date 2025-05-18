package user

import "github.com/chhongzh/chz_Base_Backend/internal/model"

func (s *Service) CreateUser(
	username string,
	password string,

	publicName string,
	personalizedSignature string,
) (*model.User, error) {
	user := &model.User{
		Username: username,
		Password: password,

		PublicName:            publicName,
		PersonalizedSignature: personalizedSignature,
	}
	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
