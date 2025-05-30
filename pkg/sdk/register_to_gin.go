package sdk

import (
	"github.com/gin-gonic/gin"
)

// 为应用程序提供一些基础的 SDK 功能
func (b *BaseSDK) RegisterToGin(r *gin.RouterGroup) {
	r.GET("/", func(ctx *gin.Context) { ctx.String(200, "OKOK!") })

	sdkGroup := r.Group("/sdk")
	{
		sdkGroup.GET("/openOAuth", b.handlerOpenOAuth)
	}
}
