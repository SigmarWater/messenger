package converter 

import(
	pb "github.com/SigmarWater/messenger/chat/pkg/api/chat_service"
	"github.com/SigmarWater/messenger/chat/internal/model"
)

func ChatDescToService(req pb.CreateRequest) *model.ChatService{
	return &model.ChatService{
		ChatName: req.GetChatName(),
	}
}

