package service

import (
	"magic-dashboard-backend/internal/controller/http/v1/request"
	"magic-dashboard-backend/internal/service/pipeline"
)

type Generation interface {
	Generate(req *request.GenerateRequest) (*pipeline.Context, error)
}

type Service struct {
	GenerationService Generation
}

func NewService() *Service {
	return &Service{
		GenerationService: NewGenerateService(),
	}
}
