package service

import (
	v1 "magic-dashboard-backend/internal/controller/http/v1"
	"strings"
)

type AuthServiceInterface interface {
	Auth(string) error
}

type EnvValidator interface {
	Validate(string) error
}

type DbTableAnalyzeServiceInterface interface {
	// GetSingularEntityNameFromTableName(string) (string, error)
	// GetTableNames(tables []v1.Table) []string
	GetEntitiesNamesFromTableNames(tables []v1.Table) []string
}

type GenerateService struct {
	AuthService           AuthServiceInterface
	PhpValidator          EnvValidator
	LaravelValidator      EnvValidator
	NovaValidator         EnvValidator
	DBMSValidator         EnvValidator
	DbTableAnalyzeService DbTableAnalyzeServiceInterface
}

func NewGenerateService(
	authService AuthServiceInterface,
	phpValidator EnvValidator,
	laravelValidator EnvValidator,
	novaValidator EnvValidator,
	dbmsValidator EnvValidator,
	dbTableAnalyzeService DbTableAnalyzeServiceInterface,
) *GenerateService {
	return &GenerateService{
		AuthService:           authService,
		PhpValidator:          phpValidator,
		LaravelValidator:      laravelValidator,
		NovaValidator:         novaValidator,
		DBMSValidator:         dbmsValidator,
		DbTableAnalyzeService: dbTableAnalyzeService,
	}
}

func (s *GenerateService) Generate(req *v1.Request) (string, error) {
	if err := s.AuthService.Auth("xxxx-xxxx-xxxx-xxxx"); err != nil {
		return "", err
	}

	if err := s.PhpValidator.Validate(req.ClientEnvironment.PhpVersion); err != nil {
		return "", err
	}

	if err := s.LaravelValidator.Validate(req.ClientEnvironment.LaravelVersion); err != nil {
		return "", err
	}

	if err := s.NovaValidator.Validate(req.ClientEnvironment.NovaVersion); err != nil {
		return "", err
	}

	if err := s.DBMSValidator.Validate(req.ClientEnvironment.DbmsName); err != nil {
		return "", err
	}

	// tableNamesList := s.DbTableAnalyzeService.GetTableNames(req.Database.Tables)
	entityNameList := s.DbTableAnalyzeService.GetEntitiesNamesFromTableNames(req.Database.Tables)

	return strings.Join(entityNameList, ","), nil
}
