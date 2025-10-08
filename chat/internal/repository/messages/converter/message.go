package converter

import (
	"github.com/SigmarWater/messenger/chat/internal/model"
	rpModel  "github.com/SigmarWater/messenger/chat/internal/repository/messages/model"
)

func ToMessageFromRepo(msg *rpModel.MessageRepository) *model.MessageService{
	return &model.MessageService{
		ChatId: msg.Id_chat,
		From_user: msg.From_user,
		Text_message: msg.Text_message,
		Time_at: msg.Time_at,
	}
}