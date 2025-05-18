package handler

import "github.com/chhongzh/chz_Base_Backend/internal/model"

func (h *Handler) Run(host string) error {
	// Migrate
	h.db.AutoMigrate(
		&model.User{},
		&model.Action{},
		&model.Application{},
		&model.Announcement{},
		&model.Permission{},
	)

	// 启动各个服务
	go h.actionService.Loop()
	go h.signService.Loop()

	h.beforeRun()

	return h.gin.Run(host)
}
