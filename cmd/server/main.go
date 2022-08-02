package main

import (
	"context"

	"github.com/kyledinh/protoserver-go/pkg/api"
	"github.com/kyledinh/protoserver-go/pkg/config"
	"github.com/kyledinh/protoserver-go/pkg/proto"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	config.LoadConfig()

	// Zap Logger, json structured
	logger := proto.Logger(ctx)

	logger.Info("MAIN")
	serverPort := viper.GetInt("serverPort")
	logger.Info("SERVER", zap.Int("serverPort", serverPort))
	logger.Info("SERVER", zap.Bool("Ready", config.IsReady()))

	go api.StartRouter(ctx, serverPort)
	select {}
}
