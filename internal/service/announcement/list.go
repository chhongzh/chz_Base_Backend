package announcement

import (
	"github.com/chhongzh/chz_Base_Backend/internal/model"
	"github.com/chhongzh/chz_Base_Backend/internal/utils"
)

func (s *Service) ListAnnouncements(applicationID string, page int) ([]*model.Announcement, error) {
	var res []*model.Announcement
	err := utils.Pagination(s.db.Where("for_application = ?", applicationID).Order("created_at DESC"), page, 15).Find(&res).Error
	return res, err
}
