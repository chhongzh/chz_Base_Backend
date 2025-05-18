package handler

import (
	"github.com/chhongzh/chz_Base_Backend/internal/constants"
	"github.com/chhongzh/chz_Base_Backend/internal/problem"
	"github.com/chhongzh/chz_Base_Backend/internal/request"
	"github.com/chhongzh/chz_Base_Backend/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v4/net"
)

func (h *Handler) getCacheInfo(c *gin.Context) {
	var req request.GetInfo
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
	// 检查权限
	if !h.permissionService.HasPermission(user.UserID, "ROOT", constants.PermissionReadInfo) {
		response.BuildResponseWithError(c, problem.ErrNoPermission)
		return
	}

	totalHit, totalStore := h.cacher.GetCacheInfo()
	response.BuildResponse(c, response.CacheInfo{
		TotalHit:   totalHit,
		TotalStore: totalStore,
	})
}

func (h *Handler) getNetworkInfo(c *gin.Context) {
	var req request.GetInfo
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
	// 检查权限
	if !h.permissionService.HasPermission(user.UserID, "ROOT", constants.PermissionReadInfo) {
		response.BuildResponseWithError(c, problem.ErrNoPermission)
		return
	}

	netStat, _ := net.IOCounters(false)

	response.BuildResponse(c, response.NetworkInfo{
		PacketSent: netStat[0].PacketsSent,
		PacketRecv: netStat[0].PacketsRecv,
		BytesSent:  netStat[0].BytesSent,
		BytesRecv:  netStat[0].BytesRecv,
	})
}
