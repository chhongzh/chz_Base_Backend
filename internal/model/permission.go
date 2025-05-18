package model

import (
	"github.com/chhongzh/chz_Base_Backend/internal/problem"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model

	ForApplicationID string
	ForUserID        string

	PermissionID  string
	PermissionKey string

	IsDeny bool
}

func (p *Permission) BeforeCreate(tx *gorm.DB) error {
	var count int64

	p.PermissionID = uuid.NewString()

	// 判断是否存在 Applicaiton ID
	tx.Model(&Application{}).
		Where("application_id = ?", p.ForApplicationID).
		Count(&count)
	if count == 0 {
		return problem.ErrApplicationNotFound
	}

	// 判断是否存在 User ID
	tx.Model(&User{}).
		Where("user_id = ?", p.ForUserID).
		Count(&count)
	if count == 0 {
		return problem.ErrUserNotFound
	}

	// 判断是否重复
	tx.Model(&Permission{}).
		Where("permission_key = ?", p.PermissionKey).
		Where("for_application_id = ?", p.ForApplicationID).
		Where("for_user_id = ?", p.ForUserID).
		Count(&count)
	if count > 0 {
		return problem.ErrPermissionAlreadyExists
	}

	return nil
}
