package permission

import "github.com/chhongzh/chz_Base_Backend/internal/model"

func (s *Service) FlagPermission(userID string, applicationID string, permissionKey string, deny bool) error {
	// 添加Permission

	return s.db.Create(&model.Permission{
		ForApplicationID: applicationID,
		ForUserID:        userID,
		PermissionKey:    permissionKey,
		IsDeny:           deny,
	}).Error
}
