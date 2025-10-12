package chat

import(
	pb "github.com/SigmarWater/messenger/chat/pkg/api/chat_service"
	"github.com/SigmarWater/messenger/chat/internal/service"
)

type Implementation struct{
	pb.UnimplementedChatApiServer
	chatSrv service.ChatService 
	msgSrv service.MessageService
}

func NewImplementation(chatSrv service.ChatService, msgSrv service.MessageService) *Implementation{
	return &Implementation{
		chatSrv: chatSrv,
		msgSrv: msgSrv,
	}
}