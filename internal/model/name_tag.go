package model

import (
    "github.com/chhongzh/chz_Base_Backend/internal/problem"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type NameTag struct {
    gorm.Model

    nameTagID string

    tag   string `validate:"min=1,max=32"`
    color string `validate:"hexcolor"`
}

type NameTagRelation struct {
    gorm.Model

    nameTagRelationID string
    nameTagID         string
    forUserID         string
}

func (n *NameTag) BeforeCreate(tx *gorm.DB) error {
    n.nameTagID = uuid.NewString()
    return nil
}

func (n *NameTagRelation) BeforeCreate(tx *gorm.DB) error {
    var count int64

    // 1. 检查 NameTag 是否存在
    tx.Model(&NameTag{}).
        Where("name_tag_id = ?", n.nameTagID).
        Count(&count)
    if count == 0 {
        return problem.ErrNameTagNotFound
    }

    // 2. 检查 User 是否存在
    tx.Model(&User{}).
        Where("user_id = ?", n.forUserID).
        Count(&count)
    if count == 0 {
        return problem.ErrUserNotFound
    }

    // 3. 检查是否重复
    tx.Model(&NameTagRelation{}).
        Where("name_tag_id = ?", n.nameTagID).
        Where("for_user_id = ?", n.forUserID).
        Count(&count)
    if count > 0 {
        return problem.ErrNameTagRelationAlreadyExists
    }

    return nil
}

func (n *NameTag) BeforeSave(tx *gorm.DB) error {
    // 检查 NameTag 的颜色
    err := validator.Struct(n)
    if err != nil {
        return err
    }
    return nil
}