package chat

import(
	"context"
	pb "github.com/SigmarWater/messenger/chat/pkg/api/chat_service"
	"github.com/SigmarWater/messenger/chat/internal/converter"
	"log"
)

func (i *Implementation) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error){
	chat, err := i.chatSrv.CreateChat(ctx, converter.ChatDescToService(req))
	
	if err != nil{
		log.Printf("Ошибка при преобразовании desc в service: %v\n", err)
		return nil, err
	}
	
	return converter.ChatServiceToDesc(chat), nil 
}