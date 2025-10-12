package users 

import(
	"context"
	"github.com/SigmarWater/messenger/auth/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64)(*model.UserService, error){
	user, err := s.userRepository.GetUser(ctx, int(id))
	
	if err != nil{
		return nil, err 
	}

	return user, err
}