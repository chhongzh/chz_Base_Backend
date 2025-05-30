package sdk

import (
	"github.com/chhongzh/chz_Base_Backend/pkg/shortcuts"
	"github.com/gin-gonic/gin"
)

func (b *BaseSDK) handlerOpenOAuth(c *gin.Context) {
	// 调用 SDK 创建
	signSessionID, err := b.OpenOAuth()
	if err != nil {
		shortcuts.BuildResponseWithError(c, err)
		return
	}

	shortcuts.BuildResponse(c, &OpenOAuthResponse{
		OAuthSessionID: signSessionID,
	})
}
