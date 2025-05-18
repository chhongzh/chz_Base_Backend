package handler

import (
	"github.com/chhongzh/chz_Base_Backend/internal/constants"
	"github.com/chhongzh/chz_Base_Backend/internal/problem"
	"github.com/chhongzh/chz_Base_Backend/internal/request"
	"github.com/chhongzh/chz_Base_Backend/internal/response"
	"github.com/chhongzh/chz_Base_Backend/internal/service/action"
	"github.com/gin-gonic/gin"
)

func (h *Handler) announcementList(c *gin.Context) {
	var req request.AnnouncementList
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

	// 检查是否有 Announcement 权限
	if !h.permissionService.HasPermission(user.UserID, "ROOT", constants.PermissionAnnouncement) {
		response.BuildResponseWithError(c, problem.ErrNoPermission)
		return
	}

	// 获取 Announcement List
	announcements, err := h.announcementService.ListAnnouncements(req.ApplicationID, req.Page)
	if err != nil {
		response.BuildResponseWithError(c, err)
		return
	}

	var announcementList []response.AnnouncementInfo
	for _, announcement := range announcements {
		announcementList = append(announcementList, response.AnnouncementInfo{
			AnnouncementID: announcement.AnnouncementID,
			Title:          announcement.Title,
			Content:        announcement.Content,
			WhoAnnounce:    announcement.WhoAnnounce,
			CreatedAt:      announcement.CreatedAt,
		})
	}

	response.BuildResponse(c, response.AnnouncementList{
		AnnouncementList: announcementList,
	})
}

func (h *Handler) announcementAnnounce(c *gin.Context) {
	var req request.AnnoucnementAnnounce
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

	// 检查是否有 Announcement 权限
	if !h.permissionService.HasPermission(user.UserID, "ROOT", constants.PermissionAnnouncement) {
		response.BuildResponseWithError(c, problem.ErrNoPermission)
		return
	}

	// 发布公告
	err = h.announcementService.Anounce(req.ApplicationID, req.Title, req.Content, user.UserID)
	if err != nil {
		response.BuildResponseWithError(c, err)
		return
	}

	// Action 记录
	h.actionService.Commit("ROOT", "[公告] 发布", action.NewMetaChain().
		WithClientInfo(c).
		Add("User ID", user.UserID).
		Add("Application ID", req.ApplicationID),
	)
}

func (h *Handler) announcementLast(c *gin.Context) {
	var req request.AnnouncementLast
	err := c.BindJSON(&req)
	if err != nil {
		return
	}

	announcement, err := h.announcementService.LastAnnouncement(req.ApplicationID)
	if err != nil {
		response.BuildResponseWithError(c, err)
		return
	}

	response.BuildResponse(c, response.AnnouncementInfo{
		AnnouncementID: announcement.AnnouncementID,
		Title:          announcement.Title,
		Content:        announcement.Content,
		WhoAnnounce:    announcement.WhoAnnounce,
		CreatedAt:      announcement.CreatedAt,
	})
}

func (h *Handler) announcementDelete(c *gin.Context) {
	var req request.AnnouncementDelete
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

	// 检查是否有 Announcement 权限
	if !h.permissionService.HasPermission(user.UserID, "ROOT", constants.PermissionAnnouncement) {
		response.BuildResponseWithError(c, problem.ErrNoPermission)
		return
	}

	// 删除公告
	err = h.announcementService.Delete(req.AnnouncementID)
	if err != nil {
		response.BuildResponseWithError(c, err)
		return
	}
}
