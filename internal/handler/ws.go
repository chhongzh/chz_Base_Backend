package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) wsOpen(c *gin.Context) {
	authToken := c.Query("authToken")
	applicationID := c.Query("applicationID")

	h.logger.Debug("Ws尝试打开", zap.String("AuthToken", authToken), zap.String("Application ID", applicationID))

	// 尝试升级
	conn, err := h.wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		h.logger.Error("Ws 升级失败", zap.Error(err))
		return
	}
	_ = conn
}
