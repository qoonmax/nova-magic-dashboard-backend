package app

import (
	v1 "magic-dashboard-backend/internal/controller/http/v1"
	"magic-dashboard-backend/internal/server/http"
	"magic-dashboard-backend/internal/service"
	"magic-dashboard-backend/internal/service/auth"
	"magic-dashboard-backend/internal/service/client_env_validate"
	"magic-dashboard-backend/internal/service/db_table_analyze"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

	phpValidator := client_env_validate.NewPhpValidator()
	laravelValidator := client_env_validate.NewLaravelValidator()
	novaValidator := client_env_validate.NewNovaValidator()
	dbmsValidator := client_env_validate.NewDBMSValidator()

	dbTableAnalyzeService := db_table_analyze.NewService()

	generateService := service.NewGenerateService(
		authService,
		phpValidator,
		laravelValidator,
		novaValidator,
		dbmsValidator,
		dbTableAnalyzeService,
	)

	handler := v1.NewHandler(generateService)
	router := v1.NewRouter(handler)
	server := new(http.Server)

	if err := server.Run(viper.GetString("port"), router.InitRoutes()); err != nil {
		logrus.Fatalf("Error server (%s)", err.Error())
	}
}
