package users

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/SigmarWater/messenger/auth/internal/model"
	"time"
)

func (s *ServiceSuite) TestCreateSuccess(){
	
	// детерминированный ожидаемый id от репозитория
	excepted_id := int64(42)
	name    := gofakeit.Name()
	email := gofakeit.Email()
	enterPassword := gofakeit.UUID()
	confirmPassword := gofakeit.UUID()
	role := "user"
	createAt := time.Now()
	updateAt := time.Now().Add(time.Hour)

	user := &model.UserService{
		Id: int(excepted_id), 
		Name: name,
		Email: email,
		EnterPassword: enterPassword,
		ConfirmPassword: confirmPassword,
		Role: role,
		CreateAt: createAt,
		UpdateAt: updateAt,
	}

	s.userRepository.On("InsertUser", s.ctx, user).Return(excepted_id, nil)
	
	id, err := s.service.Create(s.ctx, user)

	s.Require().NoError(err)
	s.Require().Equal(excepted_id, id)
}

func (s *ServiceSuite) TestCreateFail(){
	// детерминированный ожидаемый id от репозитория
	repoErr := gofakeit.Error()
	excepted_id := int64(0)
	name    := gofakeit.Name()
	email := gofakeit.Email()
	enterPassword := gofakeit.UUID()
	confirmPassword := gofakeit.UUID()
	role := "user"
	createAt := time.Now()
	updateAt := time.Now().Add(time.Hour)

	user := &model.UserService{
		Id: int(excepted_id), 
		Name: name,
		Email: email,
		EnterPassword: enterPassword,
		ConfirmPassword: confirmPassword,
		Role: role,
		CreateAt: createAt,
		UpdateAt: updateAt,
	}
	
	s.userRepository.On("InsertUser", s.ctx, user).Return(int64(0), repoErr)

	id, err := s.service.Create(s.ctx, user)

	s.Require().Error(err)
	s.Require().Equal(id, int64(0))
}