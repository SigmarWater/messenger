package converter

import (
	"log"

	"github.com/SigmarWater/messenger/chat/internal/model"
	pb "github.com/SigmarWater/messenger/chat/pkg/api/chat_service"
)

func ChatDescToService(req *pb.CreateRequest) *model.ChatService {
	chatName := req.GetChatName()
	log.Printf("DEBUG: req.GetChatName() = '%s' (type: %T)", chatName, chatName)

	return &model.ChatService{
		ChatName: chatName,
	}
}

func ChatServiceToDesc(model *model.ChatService) *pb.CreateResponse {
	return &pb.CreateResponse{
		Id:       int64(model.IdChat),
		ChatName: model.ChatName,
	}
}
