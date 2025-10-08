package memberships

import (
	"github.com/ahmadalaik/music-catalog/internal/configs"
	"github.com/ahmadalaik/music-catalog/internal/models/memberships"
)

//go:generate mockgen -source=service.go -destination=service_mock_test.go -package=memberships
type repository interface {
	CreateUser(model memberships.User) error
	GetUser(email, username string, id uint) (*memberships.User, error)
}

type service struct {
	config     *configs.Config
	repository repository
}

func NewService(config *configs.Config, repository repository) *service {
	return &service{
		config:     config,
		repository: repository,
	}
}
