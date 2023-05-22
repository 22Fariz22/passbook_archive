package main

import (
	"github.com/22Fariz22/passbook/server/config"
	"github.com/22Fariz22/passbook/server/internal/app"
	"github.com/22Fariz22/passbook/server/pkg/logger"
	"github.com/22Fariz22/passbook/server/pkg/postgres"
	"github.com/22Fariz22/passbook/server/pkg/redis"
	"github.com/22Fariz22/passbook/server/pkg/utils"
	"log"
	"os"
)

func main() {
	log.Println("Starting service")

	configPath := utils.GetConfigPath(os.Getenv("config"))
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("Loading config: %v", err)
	}

	appLogger := logger.NewAPILogger(cfg)
	appLogger.InitLogger()
	appLogger.Infof(
		"AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v",
		cfg.Server.AppVersion,
		cfg.Logger.Level,
		cfg.Server.Mode,
		cfg.Server.SSL,
	)
	appLogger.Infof("Success parsed config: %#v", cfg.Server.AppVersion)

	psqlDB, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		appLogger.Fatalf("Postgresql init: %s", err)
	}
	defer psqlDB.Close()

	redisClient := redis.NewRedisClient(cfg)
	defer redisClient.Close()
	appLogger.Info("Redis connected")

	authServer := app.NewAuthServer(appLogger, cfg, psqlDB, redisClient)
	appLogger.Fatal(authServer.Run())
}
