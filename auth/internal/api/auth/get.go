package auth

import(
	"context"
	pb "github.com/SigmarWater/messenger/auth/pkg/api/auth_service"
	"log"
	"github.com/SigmarWater/messenger/auth/internal/converter"
)

func (i *Implementation) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	user, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("Получаем user: %+#v\n", user)


	return converter.ToDescFromUser(user), nil
}