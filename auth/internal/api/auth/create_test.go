package auth

import (
	pb "github.com/SigmarWater/messenger/auth/pkg/api/auth_service"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/mock"
)

func (s *ServiceSuite) TestCreateSuccess() {
	name := gofakeit.Name()
	email := gofakeit.Email()
	enterPassword := gofakeit.UUID()
	// confirm должен совпадать с password, иначе Implementation.Create вернёт ошибку
	confirmPassword := enterPassword

	req := &pb.CreateRequest{
		Name:            name,
		Email:           email,
		Password:        enterPassword,
		PasswordConfirm: confirmPassword,
		Role:            pb.Role_ROLE_ADMIN,
	}

	s.userService.On("Create", s.ctx, mock.Anything).Return(int64(1), nil)

	resp_result, err_result := s.impl.Create(s.ctx, req)

	s.Require().NoError(err_result)
	s.Require().Equal(int64(1), resp_result.GetId())
}

func (s *ServiceSuite) TestCreateFail() {
	errTest := gofakeit.Error()
	name := gofakeit.Name()
	email := gofakeit.Email()
	enterPassword := gofakeit.UUID()
	// confirm должен совпадать с password, иначе Implementation.Create вернёт ошибку
	confirmPassword := enterPassword

	req := &pb.CreateRequest{
		Name:            name,
		Email:           email,
		Password:        enterPassword,
		PasswordConfirm: confirmPassword,
		Role:            pb.Role_ROLE_ADMIN,
	}

	s.userService.On("Create", s.ctx, mock.Anything).Return(int64(0), errTest)

	resp_result, err_result := s.impl.Create(s.ctx, req)

	s.Require().Error(err_result)
	s.Require().Equal(int64(0), resp_result.GetId())

}