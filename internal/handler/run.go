package handler

import (
	"net"

	"github.com/chhongzh/chz_Base_Backend/internal/model"
	"github.com/chhongzh/chz_Base_Backend/internal/utils"
)

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

	// 创建代理 Listener
	listener, err := net.Listen("tcp", apiHost)
	if err != nil {
		return err
	}
	listener = utils.ProxyListener(listener, h.logger.Named("Gin Listener"))

	return h.gin.RunListener(listener)
}
