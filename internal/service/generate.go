package service

import (
	"magic-dashboard-backend/internal/controller/http/v1/request"
	"magic-dashboard-backend/internal/service/pipeline"
)

type GenerateService struct {
}

func NewGenerateService() *GenerateService {
	return &GenerateService{}
}

func (s *GenerateService) Generate(req *request.GenerateRequest) (*pipeline.Context, error) {
	pctx := pipeline.NewContext(req)
	p := pipeline.NewPipeline(
		pipeline.ValidateDatabaseDriver,
		pipeline.ValidateLaravelVersion,
		pipeline.ValidateNovaVersion,
		pipeline.ValidatePhpVersion,
		pipeline.GenerateModels,
	)

	pctx, err := p.Execute(pctx)

	if err != nil {
		return nil, err
	}

	return pctx, nil
}
