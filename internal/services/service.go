package services

import (
	"magic-dashboard-backend/internal/http/requests"
	"magic-dashboard-backend/internal/services/generate_service"
)

type Service struct {
	GenerationService GenerationService
}

type GenerationService interface {
	Generate(req *requests.GenerateRequest) ([]generate_service.Instruction, error)
}

func NewService(generateService GenerationService) *Service {
	return &Service{
		GenerationService: generateService,
	}
}
