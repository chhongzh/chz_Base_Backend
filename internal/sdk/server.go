package sdk

import (
	"net"

	"github.com/chhongzh/chz_Base_Backend/internal/service/permission"
	"github.com/chhongzh/chz_Base_Backend/internal/service/security"
	"github.com/chhongzh/chz_Base_Backend/internal/service/user"
	"github.com/chhongzh/chz_Base_Backend/internal/utils"
	"github.com/chhongzh/chz_Base_Backend/pkg/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type SdkServer struct {
	pb.UnimplementedBaseSDKServiceServer

	userService       *user.Service
	securityService   *security.Service
	permissionService *permission.Service

	logger *zap.Logger
}

func NewSdkServer(logger *zap.Logger, userService *user.Service, securityService *security.Service, permissionService *permission.Service) *SdkServer {
	return &SdkServer{
		logger: logger,

		userService:       userService,
		securityService:   securityService,
		permissionService: permissionService,
	}
}

func (s *SdkServer) Run(host string) error {
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}

	// 生成一个Proxy过后的Listener
	listener = utils.ProxyListener(listener, s.logger.Named("SDK Listener"))

	grpcServer := grpc.NewServer()
	pb.RegisterBaseSDKServiceServer(grpcServer, s)

	s.logger.Info("Sdk Server 启动", zap.String("host", host))

	return grpcServer.Serve(listener)
}
