package users

import (
	"context"
	"testing"
	"time"

	"github.com/SigmarWater/messenger/auth/internal/repository/mocks"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite

	ctx context.Context

	userRepository *mocks.UserRepository

	userCache *mocks.UserCache

	service *serv
}

// Пожготовительные операции перед вызовом тестов
func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

	s.userRepository = mocks.NewUserRepository(s.T())

	s.userCache = mocks.NewUserCache(s.T())

	s.service = NewService(s.userRepository, s.userCache, time.Second)
}

// Если нам нужно подчистить после тестов
func (s *ServiceSuite) TearDownTest() {}

// Имя может быть любым
func TestServiceIntegration(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
