package chat 

import(
	"context"
	pb "github.com/SigmarWater/messenger/chat/pkg/api/chat_service"
	"github.com/SigmarWater/messenger/chat/internal/converter"
	"log"
)

func (i *Implementation) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error){
	id, err := i.msgSrv.SendMessage(ctx, converter.MessageFromDesc(req))
	if err != nil{
		log.Printf("Ошибка при отправке сообщения: %v\n", err)
		return nil, err
	}

	return &pb.SendMessageResponse{MessageID: int64(id)}, nil
}
