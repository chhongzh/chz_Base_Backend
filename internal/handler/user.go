package handler

import (
	"github.com/chhongzh/chz_Base_Backend/internal/problem"
	"github.com/chhongzh/chz_Base_Backend/internal/request"
	"github.com/chhongzh/chz_Base_Backend/internal/response"
	"github.com/chhongzh/chz_Base_Backend/internal/service/action"
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
		response.BuildResponseWithError(c, err)
		return
	}
	if currentUserCount >= h.maxUserCount {
		// Action 记录
		h.actionService.Commit("ROOT", "[用户] 用户数量达到上限", action.NewMetaChain().
			WithClientInfo(c),
		)

		response.BuildResponseWithError(c, problem.ErrUserMaxCountLimit)
		return
	}

	// 创建用户
	user, err := h.userService.CreateUser(req.Username, req.Password, req.PublicName, req.PersonalizedSignature)
	if err != nil {
		h.logger.Error(
			"用户注册失败",
			zap.Error(err),
		)

		response.BuildResponseWithError(c, err)
		return
	}

	// Action 记录
	h.actionService.Commit("ROOT", "[用户] 注册", action.NewMetaChain().
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
		response.BuildResponseWithError(c, problem.ErrLoginFailed)
		return
	}

	if user.Password != req.Password {
		response.BuildResponseWithError(c, problem.ErrLoginFailed)
		return
	}

	// 检查是否被封禁
	if user.Banned {
		response.BuildResponseWithError(c, problem.ErrUserHasBeenBanned)
		return
	}

	// 签发 Auth Token
	authToken, err := h.securityService.UserIDToAuthToken(user.UserID)
	if err != nil {
		response.BuildResponseWithError(c, err)
		return
	}

	// Action 记录
	h.actionService.Commit("ROOT", "[用户] 登录", action.NewMetaChain().
		Add("User ID", user.UserID).
		WithClientInfo(c),
	)

	response.BuildResponse(c, response.UserLogin{
		AuthToken: authToken,
	})
}
