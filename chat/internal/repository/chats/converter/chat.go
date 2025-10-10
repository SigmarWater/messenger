package converter 

import(
	"github.com/SigmarWater/messenger/chat/internal/model"
	msgRepo "github.com/SigmarWater/messenger/chat/internal/repository/chats/model"
)

func MsgRepoToService(msg msgRepo.ChatRepository) *model.ChatService{
	return &model.ChatService{
		IdChat: msg.IdChat,
		ChatName: msg.ChatName,
	}
}