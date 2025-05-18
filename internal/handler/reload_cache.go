package handler

import (
	"github.com/chhongzh/chz_Base_Backend/internal/constants"
	"github.com/chhongzh/chz_Base_Backend/internal/problem"
	"github.com/chhongzh/chz_Base_Backend/internal/request"
	"github.com/chhongzh/chz_Base_Backend/internal/response"
	"github.com/chhongzh/chz_Base_Backend/internal/service/action"
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
		response.BuildResponseWithError(c, err)
		return
	}

	// 判断是否有 Reload Cache 权限
	if !h.permissionService.HasPermission(user.UserID, "ROOT", constants.PermissionReloadCache) {
		response.BuildResponseWithError(c, problem.ErrNoPermission)
		return
	}

	// Action 记录
	h.actionService.Commit("ROOT", "[重载] 缓存", action.NewMetaChain().
		Add("User ID", user.UserID).
		WithClientInfo(c),
	)
}
