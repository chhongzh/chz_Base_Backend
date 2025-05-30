package handler

import (
	"github.com/chhongzh/chz_Base_Backend/internal/problem"
	"github.com/chhongzh/chz_Base_Backend/internal/request"
	"github.com/chhongzh/chz_Base_Backend/internal/response"
	"github.com/chhongzh/chz_Base_Backend/pkg/shortcuts"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) userRegister(c *gin.Context) {
	var req request.UserRegister
	err := c.BindJSON(&req)
	if err != nil {
		return
	}

	// 检查是否到达上限
	currentUserCount, err := h.userService.Count()
	if err != nil {
		shortcuts.BuildResponseWithError(c, err)
		return
	}
	if currentUserCount >= h.maxUserCount {
		// Action 记录
		h.actionService.Commit("ROOT", "[用户] 用户数量达到上限", shortcuts.NewMetaChain().
			WithClientInfo(c),
		)

		shortcuts.BuildResponseWithError(c, problem.ErrUserMaxCountLimit)
		return
	}

	// 创建用户
	user, err := h.userService.CreateUser(req.Username, req.Password, req.PublicName, req.PersonalizedSignature)
	if err != nil {
		h.logger.Error(
			"用户注册失败",
			zap.Error(err),
		)

		shortcuts.BuildResponseWithError(c, err)
		return
	}

	// Action 记录
	h.actionService.Commit("ROOT", "[用户] 注册", shortcuts.NewMetaChain().
		Add("User ID", user.UserID).
		WithClientInfo(c),
	)
}

func (h *Handler) userLogin(c *gin.Context) {
	var req request.UserLogin
	err := c.BindJSON(&req)
	if err != nil {
		return
	}

	user, err := h.userService.FromUsername(req.Username)
	if err != nil {
		shortcuts.BuildResponseWithError(c, problem.ErrLoginFailed)
		return
	}

	if user.Password != req.Password {
		shortcuts.BuildResponseWithError(c, problem.ErrLoginFailed)
		return
	}

	// 检查是否被封禁
	if user.Banned {
		shortcuts.BuildResponseWithError(c, problem.ErrUserHasBeenBanned)
		return
	}

	// 签发 Auth Token
	authToken, err := h.securityService.UserIDToAuthToken(user.UserID)
	if err != nil {
		shortcuts.BuildResponseWithError(c, err)
		return
	}

	// Action 记录
	h.actionService.Commit("ROOT", "[用户] 登录", shortcuts.NewMetaChain().
		Add("User ID", user.UserID).
		WithClientInfo(c),
	)

	shortcuts.BuildResponse(c, response.UserLogin{
		AuthToken: authToken,
	})
}
