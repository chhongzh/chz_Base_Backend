package main

import (
	"github.com/chhongzh/chz_Base_Backend/internal/cacher"
	"github.com/chhongzh/chz_Base_Backend/internal/handler"
	"github.com/chhongzh/chz_Base_Backend/internal/profile"
	"github.com/gin-gonic/gin"
	"github.com/go-gorm/caches/v4"
	"github.com/spf13/viper"
	prettyconsole "github.com/thessem/zap-prettyconsole"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var prof profile.Profile
var commit string

func main() {
	// 加载配置
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config.yaml")
	viper.ReadInConfig()
	viper.Unmarshal(&prof)

	logger, err := initLogger()
	if err != nil {
		panic(err)
	}

	// 打印 commit
	logger.Info("Chz Base", zap.String("commit", commit))

	// 打印配置
	logger.Info("Profile", zap.Any("prof", prof))

	// 初始化cacher
	cacher, err := initCacher(logger.Named("Cacher"))
	if err != nil {
		logger.Fatal("init cacher error", zap.Error(err))
	}

	// 初始化db
	db, err := initDb(cacher)
	if err != nil {
		logger.Fatal("init db error", zap.Error(err))
	}

	// 初始化 gin
	ginEngine, err := initGin()
	if err != nil {
		logger.Fatal("init gin error", zap.Error(err))
	}

	// 初始化 handler
	handler := handler.New(db, ginEngine, logger, commit, prof.Api.Prefix, cacher, prof.Api.MaxUserCount, prof.Api.CorsOrigins, prof.Api.MaxSignSessionCount, prof.Api.MaxSignSessionWaitingCount)

	// 运行
	err = handler.Run(prof.Api.Host, prof.Sdk.Host)
	if err != nil {
		logger.Fatal("run error", zap.Error(err))
	}
}

func initLogger() (*zap.Logger, error) {
	if !prof.IsProd {
		return prettyconsole.NewLogger(zap.DebugLevel), nil
	} else {
		return zap.NewProduction()
	}
}

func initCacher(logger *zap.Logger) (*cacher.Cacher, error) {
	return cacher.NewCacher(logger), nil
}

func initDb(cacher *cacher.Cacher) (*gorm.DB, error) {
	cacherPlugin := &caches.Caches{Conf: &caches.Config{
		Cacher: cacher,
		Easer:  true,
	}}

	db, err := gorm.Open(mysql.Open(prof.Database.Dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.Use(cacherPlugin)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func initGin() (*gin.Engine, error) {
	return gin.New(), nil
}
