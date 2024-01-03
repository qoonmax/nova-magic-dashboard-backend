package auth

import "errors"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (a *Service) Auth(key string) error {
	if key == "xxxx-xxxx-xxxx-xxxx" {
		return nil
	} else {
		return errors.New("auth key is not valid")
	}
}
