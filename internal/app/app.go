package app

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	v1 "magic-dashboard-backend/internal/controller/http/v1"
	"magic-dashboard-backend/internal/server/http"
	"magic-dashboard-backend/internal/service"
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

	serviceContainer := service.NewService()
	handler := v1.NewHandler(serviceContainer)
	router := v1.NewRouter(handler)
	server := new(http.Server)

	if err := server.Run(viper.GetString("port"), router.InitRoutes()); err != nil {
		logrus.Fatalf("Error server (%s)", err.Error())
	}
}
