package announcement

import "github.com/chhongzh/chz_Base_Backend/internal/model"

func (s *Service) Delete(announcementID string) error {
	return s.db.Delete(&model.Announcement{}, "announcement_id = ?", announcementID).Error
}
