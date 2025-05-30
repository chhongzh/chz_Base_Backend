package handler

import (
	"github.com/chhongzh/chz_Base_Backend/internal/request"
	"github.com/chhongzh/chz_Base_Backend/internal/response"
	"github.com/chhongzh/chz_Base_Backend/pkg/shortcuts"
	"github.com/gin-gonic/gin"
)

func (h *Handler) myselfPermissionList(c *gin.Context) {
	var req request.MyselfPermissionList
	err := c.BindJSON(&req)
	if err != nil {
		return
	}

	// 加载用户
	user, err := h.userFromAuthToken(req.AuthToken)
	if err != nil {
		shortcuts.BuildResponseWithError(c, err)
		return
	}

	// 获取 Permission List
	// Myself 默认获取的是 Root 的权限, 用于检查是否可以打开面板这些
	permissions, err := h.permissionService.ListPermissions(user.UserID, "ROOT")
	if err != nil {
		shortcuts.BuildResponseWithError(c, err)
		return
	}

	var permissionList []response.Permission
	for _, permission := range permissions {
		permissionList = append(permissionList, response.Permission{
			PermissionKey: permission.PermissionKey,
			IsDeny:        permission.IsDeny,
		})
	}

	shortcuts.BuildResponse(c, response.MyselfPermissionList{
		PermissionList: permissionList,
	})
}
