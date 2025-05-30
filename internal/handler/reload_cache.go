package handler

import (
	"github.com/chhongzh/chz_Base_Backend/internal/constants"
	"github.com/chhongzh/chz_Base_Backend/internal/problem"
	"github.com/chhongzh/chz_Base_Backend/internal/request"
	"github.com/chhongzh/chz_Base_Backend/pkg/shortcuts"
	"github.com/gin-gonic/gin"
)

func (h *Handler) reloadCache(c *gin.Context) {
	var req request.ReloadCache
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

	// 判断是否有 Reload Cache 权限
	if !h.permissionService.HasPermission(user.UserID, "ROOT", constants.PermissionReloadCache) {
		shortcuts.BuildResponseWithError(c, problem.ErrNoPermission)
		return
	}

	// Action 记录
	h.actionService.Commit("ROOT", "[重载] 缓存", shortcuts.NewMetaChain().
		Add("User ID", user.UserID).
		WithClientInfo(c),
	)
}
