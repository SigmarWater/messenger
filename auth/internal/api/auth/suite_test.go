package auth

import(
	"github.com/stretchr/testify/suite"
	"github.com/SigmarWater/messenger/auth/internal/service/mocks"
	"context"
	"testing"
)

type ServiceSuite struct{
	suite.Suite

	ctx context.Context 

	userService *mocks.UsersService

	impl *Implementation
}

func (s *ServiceSuite) SetupTest(){
	s.ctx = context.Background()

	s.userService = mocks.NewUsersService(s.T())

	s.impl = NewImplementation(s.userService)
}

// Если нам нужно подчистить после тестов
func(s *ServiceSuite) TearDownTest(){}

// Чтобы suite заработал у него должна быть такая функция.
//  Название все равно, главное чтобы начиналось с test
func TestServiceIntegration(t *testing.T){
	suite.Run(t, new(ServiceSuite))
}