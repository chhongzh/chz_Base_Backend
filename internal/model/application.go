package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Application struct {
	gorm.Model

	Name          string
	Desc          string
	ApplicationID string
}

func (a *Application) BeforeCreate(tx *gorm.DB) error {
	// 条件注入 uuid
	if a.ApplicationID == "" {
		a.ApplicationID = uuid.NewString()
	}
	return nil
}
