package handler

import (
	"github.com/chhongzh/chz_Base_Backend/internal/problem"
	"github.com/chhongzh/chz_Base_Backend/internal/request"
	"github.com/chhongzh/chz_Base_Backend/internal/response"
	"github.com/chhongzh/chz_Base_Backend/pkg/shortcuts"
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
		shortcuts.BuildResponseWithError(c, err)
		return
	}

	// 判断是否Application权限
	if !h.permissionService.HasPermission(user.UserID, "ROOT", "application") {
		shortcuts.BuildResponseWithError(c, problem.ErrNoPermission)
		return
	}

	// 删除应用
	err = h.applicationService.DeleteApplication(req.ApplicationID)
	if err != nil {
		shortcuts.BuildResponseWithError(c, err)
		return
	}

	// Action 记录
	h.actionService.Commit("ROOT", "[应用] 删除", shortcuts.NewMetaChain().
		Add("User ID", user.UserID).
		Add("Application ID", req.ApplicationID).
		WithClientInfo(c),
	)
}

func (h *Handler) applicationCreate(c *gin.Context) {
	var req request.ApplicationCreate
	err := c.BindJSON(&req)
	if err != nil {
		shortcuts.BuildResponseWithError(c, err)
		return
	}

	// 加载 User
	user, err := h.userFromAuthToken(req.AuthToken)
	if err != nil {
		shortcuts.BuildResponseWithError(c, err)
		return
	}

	// 判断是否Application权限
	if !h.permissionService.HasPermission(user.UserID, "ROOT", "application") {
		shortcuts.BuildResponseWithError(c, problem.ErrNoPermission)
		return
	}

	// 创建应用
	application, err := h.applicationService.CreateApplication(req.Name, req.Desc)
	if err != nil {
		shortcuts.BuildResponseWithError(c, err)
		return
	}

	// Action 记录
	h.actionService.Commit("ROOT", "[应用] 创建", shortcuts.NewMetaChain().
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
		shortcuts.BuildResponseWithError(c, err)
		return
	}

	// 判断是否Application权限
	if !h.permissionService.HasPermission(user.UserID, "ROOT", "application") {
		shortcuts.BuildResponseWithError(c, problem.ErrNoPermission)
		return
	}

	// 获取应用列表
	applications, err := h.applicationService.ListApplication(req.Page)
	if err != nil {
		shortcuts.BuildResponseWithError(c, err)
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

	shortcuts.BuildResponse(c, applicationList)
}

// 获取公开信息
func (h *Handler) applicationPublicInfo(c *gin.Context) {
	applicationID := c.Param("ApplicationID")

	// 获取 Application
	application, err := h.applicationService.GetApplicationByApplicationID(applicationID)
	if err != nil {
		shortcuts.BuildResponseWithError(c, err)
		return
	}

	shortcuts.BuildResponse(c, response.ApplicationPublicInfo{
		Name: application.Name,
		Desc: application.Desc,
	})
}
