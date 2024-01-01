package service

type VersionService struct{}

func NewVersionService() *VersionService {
	return &VersionService{}
}

func (s *VersionService) Run() bool {
	return true
}
