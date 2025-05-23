package sdk

import (
	"context"

	"github.com/chhongzh/chz_Base_Backend/pkg/pb"
	"go.uber.org/zap"
)

func (b *BaseSDK) HasPermission(accessToken string, permissionKey string) bool {
	res, err := b.client.HasPermission(context.Background(), &pb.HasPermissionRequest{AccessToken: accessToken, PermissionKey: permissionKey})
	if err != nil {
		// 日志
		b.logger.Error("Has Permission 错误", zap.Error(err))

		return false
	}

	return res.HasPermission
}
