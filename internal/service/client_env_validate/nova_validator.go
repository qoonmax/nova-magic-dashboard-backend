package client_env_validate

import (
	"errors"
	"magic-dashboard-backend/internal/service/utils"
)

type NovaValidator struct{}

func NewNovaValidator() *NovaValidator {
	return &NovaValidator{}
}

func (s *NovaValidator) Validate(novaVersion string) error {
	const minNovaVersion = "4.0.0"

	if !utils.ValidateVersion(minNovaVersion, novaVersion) {
		return errors.New("laravel nova version is not valid")
	}
	return nil
}
