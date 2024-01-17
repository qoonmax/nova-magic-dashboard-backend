package app

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"magic-dashboard-backend/internal/http"
	"magic-dashboard-backend/internal/http/handlers"
	"magic-dashboard-backend/internal/http/routers"
	"magic-dashboard-backend/internal/services"
	"magic-dashboard-backend/internal/services/generate_service"
	"os"
)

func Run() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env (%s)", err.Error())
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	service := services.NewService(
		generate_service.NewGenerateService(),
	)
	handler := handlers.NewHandler(service)
	router := routers.NewRouter(handler)
	routes := router.InitRoutes()
	server := http.NewServer(port, routes)

	if err := server.ListenAndServe(); err != nil {
		logrus.Fatalf("error server (%s)", err.Error())
	}
}
