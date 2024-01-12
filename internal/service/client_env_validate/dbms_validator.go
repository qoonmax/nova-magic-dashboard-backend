package client_env_validate

import (
	"errors"
	"fmt"
)

type DBMSValidator struct{}

func NewDBMSValidator() *DBMSValidator {
	return &DBMSValidator{}
}

func (s *DBMSValidator) Validate(dbmsName string) error {
	if dbmsName != "pgsql" {
		return errors.New(fmt.Sprintf("%s is not supported", dbmsName))
	}
	return nil
}
