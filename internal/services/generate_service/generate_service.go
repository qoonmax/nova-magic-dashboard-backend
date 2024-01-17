package generate_service

import (
	"magic-dashboard-backend/internal/http/requests"
)

type GenerateService struct{}

func NewGenerateService() *GenerateService {
	return &GenerateService{}
}

func (s *GenerateService) Generate(req *requests.GenerateRequest) ([]Instruction, error) {
	var instructions []Instruction

	if err := validateGenerateRequest(req); err != nil {
		return nil, err
	}

	models, err := generateModels(req)
	if err != nil {
		return nil, err
	}

	resources, err := generateResources(req)
	if err != nil {
		return nil, err
	}

	instructions = append(instructions, models...)
	instructions = append(instructions, resources...)
	return instructions, nil
}
