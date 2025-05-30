package sdk

import (
	"context"

	"github.com/chhongzh/chz_Base_Backend/pkg/pb"
	"go.uber.org/zap"
)

func (s *SdkServer) OpenOAuth(ctx context.Context, req *pb.OpenOAuthRequest) (*pb.OpenOAuthResponse, error) {

	// 加载 Application
	application, err := s.applicationService.GetApplicationByApplicationSecret(req.GetApplicationSecret())
	if err != nil {
		s.logger.Error("Failed to get application by secret", zap.Stringer("Req", req), zap.Error(err))

		return nil, err
	}

	// 打开
	signSessionID, err := s.signService.CreateLoginForApplicationID(application.ApplicationID)
	if err != nil {
		s.logger.Error("Failed to create sign session", zap.Stringer("Req", req), zap.Error(err))

		return nil, err
	}

	s.logger.Info("OAuth 申请已经创建.", zap.String("SignSessionID", signSessionID), zap.String("ApplicationID", application.ApplicationID))

	return &pb.OpenOAuthResponse{
		SignSessionID: signSessionID,
	}, nil
}
