package nametag

import "github.com/chhongzh/chz_Base_Backend/internal/model"

func (s *Service) CreateNameTag(tag string, desc string, color string) (*model.NameTag, error) {
	nameTag := &model.NameTag{
		Tag:   tag,
		Desc:  desc,
		Color: color,
	}

	if err := s.db.Create(nameTag).Error; err != nil {
		return nil, err
	}

	return nameTag, nil
}
