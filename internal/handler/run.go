package handler

import "github.com/chhongzh/chz_Base_Backend/internal/model"

func (h *Handler) Run(apiHost string, sdkServerHost string) error {
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

	// 启动 SdkServer
	go h.sdkServer.Run(sdkServerHost)

	h.beforeRun()

	return h.gin.Run(apiHost)
}
