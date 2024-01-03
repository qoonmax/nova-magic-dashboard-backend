package env_validate

import (
	"errors"
	"magic-dashboard-backend/internal/service/utils"
)

type PhpValidator struct{}

func NewPhpValidator() *PhpValidator {
	return &PhpValidator{}
}

func (s *PhpValidator) Validate(phpVersion string) error {
	const minPhpVersion = "7.4.0"

	if !utils.ValidateVersion(minPhpVersion, phpVersion) {
		return errors.New("php version is not valid")
	}
	return nil
}
