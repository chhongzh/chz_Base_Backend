package handler

import (
	"github.com/chhongzh/chz_Base_Backend/internal/problem"
	"github.com/chhongzh/chz_Base_Backend/internal/request"
	"github.com/chhongzh/chz_Base_Backend/internal/response"
	"github.com/chhongzh/chz_Base_Backend/pkg/shortcuts"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) signCreate(c *gin.Context) {
	var req request.SignCreate
	err := c.BindJSON(&req)
	if err != nil {
		return
	}

	// 调用 service 创建
	signSessionID, err := h.signService.CreateLoginForApplicationID(req.ApplicationID)
	if err != nil {
		shortcuts.BuildResponseWithError(c, err)
		return
	}

	shortcuts.BuildResponse(c, response.SignCreate{
		SignSessionID: signSessionID,
	})
}

func (h *Handler) signComplete(c *gin.Context) {
	var req request.SignComplete
	err := c.BindJSON(&req)
	if err != nil {
		return
	}

	// 加载 SignSessionID
	signSessionID := c.Param("SignSessionID")

	// 加载用户
	user, err := h.userFromAuthToken(req.AuthToken)
	if err != nil {
		shortcuts.BuildResponseWithError(c, err)
		return
	}
	if user.Banned {
		shortcuts.BuildResponseWithError(c, problem.ErrUserHasBeenBanned)
		return
	}

	// 加载 ApplicationID
	signSession, err := h.signService.GetSignSession(signSessionID)
	if err != nil {
		shortcuts.BuildResponseWithError(c, err)
		return
	}

	h.logger.Info("完成 SignSession", zap.String("SignSessionID", signSessionID), zap.String("ApplicationID", signSession.ApplicationID))

	// 签发 AccessToken
	accessToken, err := h.securityService.ApplicationIDAndUserIDToAccessToken(signSession.ApplicationID, user.UserID)
	if err != nil {
		shortcuts.BuildResponseWithError(c, err)
		return
	}

	// 调用 service 完成
	err = h.signService.CompleteSignSession(signSession.SignSessionID, accessToken, nil)
	if err != nil {
		shortcuts.BuildResponseWithError(c, err)
		return
	}

	// Action 记录
	h.actionService.Commit("ROOT", "[授权登录] Oauth 授权登录完成", shortcuts.NewMetaChain().
		Add("ApplicationID", signSession.ApplicationID).
		Add("User ID", user.UserID).
		WithClientInfo(c),
	)

	// 这里上面调用的CompleteSignSession会通知 waiter 返回结果.
}

func (h *Handler) signWait(c *gin.Context) {
	// 获取 SignSessionID
	signSessionID := c.Param("SignSessionID")

	accessToken, err := h.signService.WaitForSignSession(signSessionID)
	if err != nil {
		shortcuts.BuildResponseWithError(c, err)
		return
	}

	// 调用 service 等待
	shortcuts.BuildResponse(c, response.SignWait{AccessToken: accessToken})
}

func (h *Handler) signInfo(c *gin.Context) {
	// 获取 SignSessionID
	signSessionID := c.Param("SignSessionID")

	// 获取Sign实例
	signSession, err := h.signService.GetSignSession(signSessionID)
	if err != nil {
		shortcuts.BuildResponseWithError(c, err)
		return
	}

	// 获取 Appliction 实例
	application, err := h.applicationService.GetApplicationByApplicationID(signSession.ApplicationID)
	if err != nil {
		shortcuts.BuildResponseWithError(c, err)
		return
	}

	// 返回信息
	shortcuts.BuildResponse(c, response.SignInfo{
		ApplicationName: application.Name,
		ApplicationDesc: application.Desc,
	})
}
