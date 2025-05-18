package handler

import (
	"time"

	"github.com/chhongzh/chz_Base_Backend/internal/service/action"
)

func (h *Handler) beforeRun() {
	// 检查是否有 ROOT 的 APP, 如果没有, 则创建
	_, err := h.applicationService.FromApplicationID("ROOT")
	if err != nil {
		// 创建 ROOT APP

		h.applicationService.CreateApplicationWithApplicationID(
			"ROOT",
			"核心应用, 不可以删除!",
			"ROOT",
		)
	}

	// 发送启动 Action
	h.actionService.Commit("ROOT", "[Core] 服务启动",
		action.NewMetaChain().
			Add("Start At", time.Now()).
			Add("Commit", h.commit),
	)
}
