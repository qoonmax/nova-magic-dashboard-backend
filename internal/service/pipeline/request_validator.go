package pipeline

import (
	"errors"
	"fmt"
	"magic-dashboard-backend/internal/service/utils"
)

func ValidateDatabaseDriver(pctx *Context) error {
	const databaseDriverName = "pgsql"

	if pctx.Req.ClientEnvironment.DatabaseDriverName != databaseDriverName {
		return errors.New(fmt.Sprintf("%s is not supported", pctx.Req.ClientEnvironment.DatabaseDriverName))
	}
	return nil
}

func ValidateLaravelVersion(pctx *Context) error {
	const minLaravelVersion = "9.0.0"

	if !utils.ValidateVersion(minLaravelVersion, pctx.Req.ClientEnvironment.LaravelVersion) {
		return errors.New("laravel version is not valid")
	}
	return nil
}

func ValidateNovaVersion(pctx *Context) error {
	const minNovaVersion = "4.0.0"

	if !utils.ValidateVersion(minNovaVersion, pctx.Req.ClientEnvironment.NovaVersion) {
		return errors.New("laravel nova version is not valid")
	}
	return nil
}

func ValidatePhpVersion(pctx *Context) error {
	const minPhpVersion = "7.4.0"

	if !utils.ValidateVersion(minPhpVersion, pctx.Req.ClientEnvironment.PhpVersion) {
		return errors.New("php version is not valid")
	}
	return nil
}
