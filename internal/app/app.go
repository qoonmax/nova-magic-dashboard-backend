package app

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	v1 "magic-dashboard-backend/internal/controller/http/v1"
	"magic-dashboard-backend/internal/server/http"
	"magic-dashboard-backend/internal/service"
	"magic-dashboard-backend/internal/service/auth"
	"magic-dashboard-backend/internal/service/env_validate"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func Run() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error init config (%s)", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env(%s)", err.Error())
	}

	authService := auth.NewService()
	phpValidator := env_validate.NewPhpValidator()
	laravelValidator := env_validate.NewLaravelValidator()
	novaValidator := env_validate.NewNovaValidator()
	dbmsValidator := env_validate.NewDBMSValidator()

	generateService := service.NewGenerateService(
		authService,
		phpValidator,
		laravelValidator,
		novaValidator,
		dbmsValidator,
	)

	handler := v1.NewHandler(generateService)
	router := v1.NewRouter(handler)
	server := new(http.Server)

	if err := server.Run(viper.GetString("port"), router.InitRoutes()); err != nil {
		logrus.Fatalf("Error server (%s)", err.Error())
	}
}
