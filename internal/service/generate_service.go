package service

type AuthServiceInterface interface {
	Auth(string) error
}

type EnvValidator interface {
	Validate(string) error
}

type GenerateService struct {
	AuthService      AuthServiceInterface
	PhpValidator     EnvValidator
	LaravelValidator EnvValidator
	NovaValidator    EnvValidator
	DBMSValidator    EnvValidator
}

func NewGenerateService(
	authService AuthServiceInterface,
	phpValidator EnvValidator,
	laravelValidator EnvValidator,
	novaValidator EnvValidator,
	dbmsValidator EnvValidator,
) *GenerateService {
	return &GenerateService{
		AuthService:      authService,
		PhpValidator:     phpValidator,
		LaravelValidator: laravelValidator,
		NovaValidator:    novaValidator,
		DBMSValidator:    dbmsValidator,
	}
}

func (s *GenerateService) Generate() (string, error) {
	if err := s.AuthService.Auth("xxxx-xxxx-xxxx-xxxx"); err != nil {
		return "", err
	}

	if err := s.PhpValidator.Validate("7.4.0"); err != nil {
		return "", err
	}

	if err := s.LaravelValidator.Validate("7.0.0"); err != nil {
		return "", err
	}

	if err := s.NovaValidator.Validate("4.0.0"); err != nil {
		return "", err
	}

	if err := s.DBMSValidator.Validate("pgsql"); err != nil {
		return "", err
	}

	return "Result from Generate Service", nil
}
