package sdk

import (
	"context"

	"github.com/chhongzh/chz_Base_Backend/pkg/pb"
	"go.uber.org/zap"
)

func (s *SdkServer) VerifyApplicationSecret(ctx context.Context, req *pb.VerifyApplicationSecretRequest) (*pb.VerifyApplicationSecretResponse, error) {
	// 加载
	_, err := s.applicationService.GetApplicationByApplicationSecret(req.ApplicationSecret)
	if err != nil {
		s.logger.Error("VerifyApplicationSecret: 获取应用失败", zap.Error(err))

		return &pb.VerifyApplicationSecretResponse{
			IsValid: false,
		}, nil
	}

	return &pb.VerifyApplicationSecretResponse{
		IsValid: true,
	}, nil
}
