package repository 

import "github.com/SigmarWater/messenger/chat/internal/model"

import(
	"context"
)

type MessageRepository interface{
	GetMessage(ctx context.Context, id_message int) (*model.MessageService, error)
	SendMessage(ctx context.Context, msg *model.MessageService) (id int, err error)
}


type ChatRepository interface{
	CreateChat(ctx context.Context, chatInfo *model.ChatService) (*model.ChatService, error)
	DeleteChat(ctx context.Context, idChat int64) error
}