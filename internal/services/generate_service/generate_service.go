package generate_service

import (
	"magic-dashboard-backend/internal/http/requests"
)

type GenerateService struct{}

func NewGenerateService() *GenerateService {
	return &GenerateService{}
}

func (s *GenerateService) Generate(req *requests.GenerateRequest) ([]Instruction, error) {
	if err := validateGenerateRequest(req); err != nil {
		return nil, err
	}

	models, err := generateModels(req)
	if err != nil {
		return nil, err
	}
	return models, nil
}
