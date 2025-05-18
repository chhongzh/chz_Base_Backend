package announcement

import "github.com/chhongzh/chz_Base_Backend/internal/model"

func (s *Service) LastAnnouncement(applicationID string) (*model.Announcement, error) {
	var res model.Announcement
	return &res, s.db.Where("for_application = ?", applicationID).Order("created_at DESC").First(&res).Error
}
