package nametag

import "github.com/chhongzh/chz_Base_Backend/internal/model"

func (s *Service) DeleteNameTagForUser(nameTagID string, userID string) error {
	return s.db.Where("name_tag_id = ?", nameTagID).
		Where("for_user_id = ?", userID).
		Delete(&model.NameTagOwned{}).Error
}

func (s *Service) DeleteNameTag(nameTagID string) error {
	return s.db.Where("name_tag_id = ?", nameTagID).
		Delete(&model.NameTag{}).Error
}
