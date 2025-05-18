package announcement

import "github.com/chhongzh/chz_Base_Backend/internal/model"

func (s *Service) Anounce(applicationID string, title string, content string, whoAnnounce string) error {
	return s.db.Create(&model.Announcement{
		ForApplication: applicationID,
		Title:          title,
		Content:        content,
		WhoAnnounce:    whoAnnounce,
	}).Error
}
