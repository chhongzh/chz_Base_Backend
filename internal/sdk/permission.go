package sdk

import (
	"context"

	"github.com/chhongzh/chz_Base_Backend/pkg/pb"
	"go.uber.org/zap"
)

func (s *SdkServer) HasPermission(ctx context.Context, req *pb.HasPermissionRequest) (*pb.HasPermissionResponse, error) {
	// 从AccessToken中加载用户信息
	applicationID, userID, err := s.securityService.ApplicationIDAndUserIDFromAccessToken(req.GetAccessToken())
	if err != nil {
		s.logger.Error("Sdk Has Permission Failed.", zap.Any("Request", req), zap.Error(err))
		return nil, err
	}

	// 检查权限
	has := s.permissionService.HasPermission(userID, applicationID, req.GetPermissionKey())

	return &pb.HasPermissionResponse{HasPermission: has}, nil
}
