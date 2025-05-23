package main

import (
	"log"
	"time"

	"github.com/chhongzh/chz_Base_Backend/pkg/sdk"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	sdkClient, err := sdk.NewBaseSDK(logger)
	if err != nil {
		panic(err)
	}

	go test_long_connection(sdkClient)

	select {}

	// logger.Info("Test Called", zap.Bool("Has", sdkClient.HasPermission("1123", "123")))
}

func test_long_connection(client *sdk.BaseSDK) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		log.Printf("Result Result %t", client.HasPermission("1234", "1234"))

		<-ticker.C
	}
}
