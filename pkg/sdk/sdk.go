package sdk

import (
	"github.com/chhongzh/chz_Base_Backend/pkg/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BaseSDK struct {
	client pb.BaseSDKServiceClient
	logger *zap.Logger
}

func NewBaseSDK(logger *zap.Logger) (*BaseSDK, error) {
	// 创建 Protobuf RPC 客户端
	grpcClient, err := grpc.NewClient(SDK_URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// 创建 Protobuf RPC 服务
	serviceClient := pb.NewBaseSDKServiceClient(grpcClient)

	logger.Info("SDK Connected")

	return &BaseSDK{
		client: serviceClient,
		logger: logger,
	}, nil
}
