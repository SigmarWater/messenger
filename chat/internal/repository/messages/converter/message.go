package converter

import (
	"github.com/SigmarWater/messenger/chat/internal/model"
	rpModel  "github.com/SigmarWater/messenger/chat/internal/repository/messages/model"
)

func ToMessageFromRepo(msg *rpModel.MessageRepository) *model.MessageService{
	return &model.MessageService{
		IdMessage: msg.IdMessage,
		ChatId: msg.IdChat,
		ChatName: msg.ChatName,
		FromUser: msg.FromUser,
		TextMessage: msg.TextMessage,
		TimeAt: msg.TimeAt.Time,
	}
}