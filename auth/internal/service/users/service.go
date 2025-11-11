package users

import (
	//pb "github.com/SigmarWater/messenger/auth/pkg/api/auth_service"
	"github.com/SigmarWater/messenger/auth/internal/repository"
	"time"
)

type serv struct {
	userRepository  repository.UserRepository
	cacheRepository repository.UserCache
	cacheTTL        time.Duration
}

func NewService(userRepository repository.UserRepository,
	cacheRepository repository.UserCache,
	cacheTTL time.Duration) *serv {
	return &serv{
		userRepository:  userRepository,
		cacheRepository: cacheRepository,
		cacheTTL:        cacheTTL,
	}
}
