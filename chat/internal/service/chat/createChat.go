package chat

import(
	"github.com/SigmarWater/messenger/chat/internal/model"
	"context"
)


func (s *servChat) CreateChat(ctx context.Context, chatInfo *model.ChatService) (*model.ChatService, error){
	chat, err := s.chatRepo.CreateChat(ctx, chatInfo)
	if err != nil{
		return nil, err
	}
	return chat, nil
}