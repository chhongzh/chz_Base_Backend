package model

import (
	"github.com/chhongzh/chz_Base_Backend/internal/problem"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Announcement struct {
	gorm.Model

	Title   string
	Content string

	AnnouncementID string
	WhoAnnounce    string
	ForApplication string
}

func (a *Announcement) BeforeCreate(tx *gorm.DB) error {
	// 注入 uuid
	a.AnnouncementID = uuid.NewString()

	// 检查是否存在用户
	var count int64
	tx.Model(&User{}).Where("user_id = ?", a.WhoAnnounce).Count(&count)
	if count == 0 {
		return problem.ErrUserNotFound
	}

	return nil
}
