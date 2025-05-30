package sdk

import (
	"context"

	"github.com/chhongzh/chz_Base_Backend/pkg/pb"
	"go.uber.org/zap"
)

// 打开一个 OAuth 授权对话
// 返回会话 ID
func (b *BaseSDK) OpenOAuth() (string, error) {
	res, err := b.client.OpenOAuth(context.Background(), &pb.OpenOAuthRequest{
		ApplicationSecret: b.applicationSecret,
	})
	if err != nil {
		b.logger.Error("Open OAuth 错误", zap.Error(err))

		return "", err
	}

	return res.GetSignSessionID(), nil
}
