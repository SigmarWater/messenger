package repository 

import "github.com/SigmarWater/messenger/chat/internal/model"

import(
	"context"
)

type MessageRepository interface{
	GetMessage(ctx context.Context, id_message int) (*model.MessageService, error)
	SendMessage(ctx context.Context, msg *model.MessageService) (int, error)
}


type ChatRepository interface{
	GetInfoChat(ctx context.Context, id_chat int) (*model.ChatService, error)
}