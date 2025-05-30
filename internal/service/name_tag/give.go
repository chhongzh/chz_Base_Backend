package nametag

import "github.com/chhongzh/chz_Base_Backend/internal/model"

func (s *Service) GiveNameTagToUser(nameTagID string, userID string) error {
	return s.db.Create(&model.NameTagOwned{
		NameTagID: nameTagID,
		ForUserID: userID,
	}).Error
}
