package env_validate

import (
	"errors"
	"magic-dashboard-backend/internal/service/utils"
)

type LaravelValidator struct{}

func NewLaravelValidator() *LaravelValidator {
	return &LaravelValidator{}
}

func (s *LaravelValidator) Validate(laravelVersion string) error {
	const minLaravelVersion = "9.0.0"

	if !utils.ValidateVersion(minLaravelVersion, laravelVersion) {
		return errors.New("laravel version is not valid")
	}
	return nil
}
