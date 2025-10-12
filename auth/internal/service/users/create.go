package users 

import(
	"context"
	"github.com/SigmarWater/messenger/auth/internal/model"
)

func (s *serv) Create(ctx context.Context, user *model.UserService)(int64, error){
	id, err := s.userRepository.InsertUser(ctx, user) 

	if err != nil{
		return 0, err
	}
	return int64(id) , nil
}