package model

import (
	"github.com/chhongzh/chz_Base_Backend/internal/problem"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	PublicName            string `validate:"min=1,max=64"`
	PersonalizedSignature string `validate:"max=512"`

	Username string `validate:"min=1,max=32"`
	Password string `validate:"min=1,max=256"`

	Banned bool

	UserID string
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	// 注入 uuid
	u.UserID = uuid.NewString()
	return nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	// 判断是否重复
	var count int64
	tx.Model(&User{}).Where("username = ?", u.Username).Count(&count)
	if count > 0 {
		return problem.ErrUserAlreadyExists
	}

	// 校验
	err := validator.Struct(u)
	if err != nil {
		return err
	}

	return nil
}
