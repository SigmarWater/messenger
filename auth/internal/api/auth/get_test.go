package auth

import (
	"github.com/SigmarWater/messenger/auth/internal/model"
	pb "github.com/SigmarWater/messenger/auth/pkg/api/auth_service"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/SigmarWater/messenger/auth/internal/converter"
)

func (s *ServiceSuite) TestGetSuccess() {
	req := &pb.GetRequest{
		Id: gofakeit.Int64(),
	}

	model := &model.UserService{
		Id: gofakeit.Int(),
		Name: gofakeit.Name(),
		Email: gofakeit.Email(),
		EnterPassword: gofakeit.UUID(),
		ConfirmPassword: gofakeit.UUID(),
		Role: "user",
		CreateAt: gofakeit.Date(),
		UpdateAt: gofakeit.Date(),
	}

	s.userService.On("Get", s.ctx, req.GetId()).Return(model, nil)
	resp, err := s.impl.Get(s.ctx, req)

	s.Require().Equal(converter.ToDescFromUser(model), resp)
	s.Require().NoError(err)
}

func (s *ServiceSuite) TestGetFail() {
	req := &pb.GetRequest{
		Id: gofakeit.Int64(),
	}
	errFail := gofakeit.Error()

	
	s.userService.On("Get", s.ctx, req.GetId()).Return(nil, errFail)
	resp, err := s.impl.Get(s.ctx, req)

	s.Require().Nil(resp)
	s.Require().Error(err)
}