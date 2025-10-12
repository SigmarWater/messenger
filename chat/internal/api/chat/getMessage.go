package chat

import(
	"context"
	pb "github.com/SigmarWater/messenger/chat/pkg/api/chat_service"
	"log"
	"github.com/SigmarWater/messenger/chat/internal/converter"
)

func (i *Implementation)GetMessage(ctx context.Context, req *pb.GetMessageRequest) (*pb.GetMessageResponse, error){
	msg, err := i.msgSrv.GetMessage(ctx, int(req.GetId()))
	if err != nil{
		log.Printf("Ошибка при получении сообщения: %v\n", err)
		return nil, err
	}

	return converter.MessageFromService(msg), nil
}