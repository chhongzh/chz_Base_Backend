package permission

import (
	"github.com/chhongzh/chz_Base_Backend/internal/model"
)

func (s *Service) HasPermission(userID string, applicationID string, permissionKey string) bool {
	var permission model.Permission

	err := s.db.Where("for_user_id = ?", userID).
		Where("for_application_id = ?", applicationID).
		Where("permission_key = ?", permissionKey).
		Last(&permission).Error

	if err != nil {
		// 尝试匹配通配符
		err := s.db.Where("for_user_id = ?", userID).
			Where("for_application_id = ?", applicationID).
			Where("permission_key = ?", "*").
			Last(&permission).Error

		if err != nil {
			return false
		}

		return !permission.IsDeny
	}

	// 如果是 Deny 那就拒绝
	return !permission.IsDeny
}
