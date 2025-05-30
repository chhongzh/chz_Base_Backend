package nametag

import "github.com/chhongzh/chz_Base_Backend/internal/model"

func (s *Service) ListOwnedNameTagsForUser(userID string) ([]*model.NameTagOwned, error) {
	var nameTags []*model.NameTagOwned
	err := s.db.Where("for_user_id = ?", userID).Find(&nameTags).Error
	if err != nil {
		return nil, err
	}
	return nameTags, nil
}

func (s *Service) ListAllNameTags() ([]*model.NameTag, error) {
	var nameTags []*model.NameTag
	err := s.db.Find(&nameTags).Error
	if err != nil {
		return nil, err
	}
	return nameTags, nil
}
