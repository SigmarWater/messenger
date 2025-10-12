package chat

import(
	"github.com/SigmarWater/messenger/chat/internal/repository"
	"github.com/SigmarWater/messenger/chat/internal/service"
)

type servChat struct{
	chatRepo repository.ChatRepository
}

func NewServChat(chatRepo repository.ChatRepository) service.ChatService{
	return &servChat{chatRepo: chatRepo}
}