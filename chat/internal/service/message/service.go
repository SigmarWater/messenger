package message

import(
	"github.com/SigmarWater/messenger/chat/internal/repository"
	"github.com/SigmarWater/messenger/chat/internal/service"
)

type servMessage struct{
	repoMessage repository.MessageRepository
}

func NewServMessage(repoMessage repository.MessageRepository) service.MessageService{
	return &servMessage{repoMessage: repoMessage}
}