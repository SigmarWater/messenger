package converter 

import(
	pb "github.com/SigmarWater/messenger/chat/pkg/api/chat_service"
	"github.com/SigmarWater/messenger/chat/internal/model"
)

func ChatDescToService(req *pb.CreateRequest) *model.ChatService{
	return &model.ChatService{
		ChatName: req.GetChatName(),
	}
}

func ChatServiceToDesc(model *model.ChatService) *pb.CreateResponse{
	return &pb.CreateResponse{
		Id: int64(model.IdChat),
		ChatName: model.ChatName,
	}
}
