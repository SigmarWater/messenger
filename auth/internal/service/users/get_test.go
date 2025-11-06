package users

import(
	"github.com/brianvoe/gofakeit/v7"
	"time"
	"github.com/SigmarWater/messenger/auth/internal/model"
) 

func (s *ServiceSuite) GetUserSuccess() {
	id := gofakeit.Int64()
	name    := gofakeit.Name()
	email := gofakeit.Email()
	enterPassword := gofakeit.UUID()
	confirmPassword := gofakeit.UUID()
	role := "user"
	createAt := time.Now()
	updateAt := time.Now().Add(time.Hour)

	user := &model.UserService{
		Id: int(id), 
		Name: name,
		Email: email,
		EnterPassword: enterPassword,
		ConfirmPassword: confirmPassword,
		Role: role,
		CreateAt: createAt,
		UpdateAt: updateAt,
	}

	s.userRepository.On("GetUser", s.ctx, id).Return(user,nil)
	user_result, err_result := s.service.Get(s.ctx, id)

	s.Require().NoError(err_result)
	s.Require().Equal(user, user_result)
}