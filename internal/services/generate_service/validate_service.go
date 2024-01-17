package generate_service

import (
	"errors"
	"fmt"
	"magic-dashboard-backend/internal/http/requests"
	"magic-dashboard-backend/internal/services/utils"
)

func validateGenerateRequest(req *requests.GenerateRequest) error {
	var err error

	if err = validateDatabaseDriver(req); err != nil {
		return err
	}

	if err = validateLaravelVersion(req); err != nil {
		return err
	}

	if err = validateNovaVersion(req); err != nil {
		return err
	}

	if err = validatePhpVersion(req); err != nil {
		return err
	}

	return nil
}

func validateDatabaseDriver(req *requests.GenerateRequest) error {
	const databaseDriverName = "pgsql"

	if req.ClientEnvironment.DatabaseDriverName != databaseDriverName {
		return errors.New(fmt.Sprintf("%s is not supported", req.ClientEnvironment.DatabaseDriverName))
	}
	return nil
}

func validateLaravelVersion(req *requests.GenerateRequest) error {
	const minLaravelVersion = "9.0.0"

	if !utils.ValidateVersion(minLaravelVersion, req.ClientEnvironment.LaravelVersion) {
		return errors.New("laravel version is not valid")
	}
	return nil
}

func validateNovaVersion(req *requests.GenerateRequest) error {
	const minNovaVersion = "4.0.0"

	if !utils.ValidateVersion(minNovaVersion, req.ClientEnvironment.NovaVersion) {
		return errors.New("laravel nova version is not valid")
	}
	return nil
}

func validatePhpVersion(req *requests.GenerateRequest) error {
	const minPhpVersion = "7.4.0"

	if !utils.ValidateVersion(minPhpVersion, req.ClientEnvironment.PhpVersion) {
		return errors.New("php version is not valid")
	}
	return nil
}
