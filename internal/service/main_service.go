package service

import "errors"

type VersionServiceInterface interface {
	Run() bool
}

type MainService struct {
	VersionService VersionServiceInterface
}

func NewMainService(versionService VersionServiceInterface) *MainService {
	return &MainService{
		VersionService: versionService,
	}
}

func (s *MainService) Run() (string, error) {
	if !s.VersionService.Run() {
		return "", errors.New("error in VersionService")
	}
	return "Result from MainService", nil
}
