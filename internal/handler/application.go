package handler

import (
	"github.com/chhongzh/chz_Base_Backend/internal/problem"
	"github.com/chhongzh/chz_Base_Backend/internal/request"
	"github.com/chhongzh/chz_Base_Backend/internal/response"
	"github.com/chhongzh/chz_Base_Backend/internal/service/action"
	"github.com/gin-gonic/gin"
)

func (h *Handler) applicationDelete(c *gin.Context) {
	var req request.ApplicationDelete
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

	// 判断是否Application权限
	if !h.permissionService.HasPermission(user.UserID, "ROOT", "application") {
		response.BuildResponseWithError(c, problem.ErrNoPermission)
	}

	// 删除应用
	err = h.applicationService.DeleteApplication(req.ApplicationID)
	if err != nil {
		response.BuildResponseWithError(c, err)
		return
	}

	// Action 记录
	h.actionService.Commit("ROOT", "[应用] 删除", action.NewMetaChain().
		Add("User ID", user.UserID).
		Add("Application ID", req.ApplicationID).
		WithClientInfo(c),
	)
}

func (h *Handler) applicationCreate(c *gin.Context) {
	var req request.ApplicationCreate
	err := c.BindJSON(&req)
	if err != nil {
		return
	}

	// 加载 User
	user, err := h.userFromAuthToken(req.AuthToken)
	if err != nil {
		response.BuildResponseWithError(c, err)
		return
	}

	// 判断是否Application权限
	if !h.permissionService.HasPermission(user.UserID, "ROOT", "application") {
		response.BuildResponseWithError(c, problem.ErrNoPermission)
	}

	// 创建应用
	application, err := h.applicationService.CreateApplication(req.Name, req.Desc)
	if err != nil {
		response.BuildResponseWithError(c, err)
		return
	}

	// Action 记录
	h.actionService.Commit("ROOT", "[应用] 创建", action.NewMetaChain().
		Add("User ID", user.UserID).
		Add("Application ID", application.ApplicationID).
		WithClientInfo(c),
	)
}

func (h *Handler) applicationList(c *gin.Context) {
	var req request.ApplicationList
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

	// 判断是否Application权限
	if !h.permissionService.HasPermission(user.UserID, "ROOT", "application") {
		response.BuildResponseWithError(c, problem.ErrNoPermission)
	}

	// 获取应用列表
	applications, err := h.applicationService.ListApplication(req.Page)
	if err != nil {
		response.BuildResponseWithError(c, err)
		return
	}

	var applicationList []response.ApplicationInfo
	for _, application := range applications {
		applicationList = append(applicationList, response.ApplicationInfo{
			ApplicationID: application.ApplicationID,
			Name:          application.Name,
			Desc:          application.Desc,
		})
	}

	response.BuildResponse(c, applicationList)
}
