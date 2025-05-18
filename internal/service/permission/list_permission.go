package permission

import "github.com/chhongzh/chz_Base_Backend/internal/model"

func (s *Service) ListPermissions(userID string, applicationID string) ([]*model.Permission, error) {
	var permissions []*model.Permission

	err := s.db.Where("for_user_id = ?", userID).
		Where("for_application_id = ?", applicationID).
		Find(&permissions).Error
	if err != nil {
		return nil, err
	}

	return permissions, nil
}
