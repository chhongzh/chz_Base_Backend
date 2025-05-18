package model

import (
	"github.com/chhongzh/chz_Base_Backend/internal/problem"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Action struct {
	gorm.Model

	FromApplicationID string
	Message           string
	Meta              any `gorm:"serializer:json"`

	ActionID string
}

func (a *Action) BeforeCreate(tx *gorm.DB) error {
	// 检查 Application 是否存在
	var count int64
	tx.Model(&Application{}).Where("application_id = ?", a.FromApplicationID).Count(&count)
	if count == 0 {
		return problem.ErrApplicationNotFound
	}

	// 注入 uuid
	a.ActionID = uuid.NewString()

	return nil
}
