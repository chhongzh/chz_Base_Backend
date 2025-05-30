package sdk

import (
	"context"
	"fmt"

	"github.com/chhongzh/chz_Base_Backend/pkg/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BaseSDK struct {
	applicationSecret string
	client            pb.BaseSDKServiceClient
	logger            *zap.Logger
}

func NewBaseSDK(applicationSecret string, logger *zap.Logger) (*BaseSDK, error) {
	// 创建 Protobuf RPC 客户端
	grpcClient, err := grpc.NewClient(SDK_URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	// 创建 Protobuf RPC 服务
	serviceClient := pb.NewBaseSDKServiceClient(grpcClient)

	// 验证密钥是否正确
	res, err := serviceClient.VerifyApplicationSecret(context.Background(), &pb.VerifyApplicationSecretRequest{ApplicationSecret: applicationSecret})
	if err != nil {
		return nil, err
	}
	if !res.IsValid {
		return nil, fmt.Errorf("invalid application secret: %s", applicationSecret)
	}

	logger.Info("SDK Connected")

	return &BaseSDK{
		applicationSecret: applicationSecret,
		client:            serviceClient,
		logger:            logger,
	}, nil
}
