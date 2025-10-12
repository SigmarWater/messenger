package message

import(
	"context"
	"github.com/SigmarWater/messenger/chat/internal/model"
)

func (s *servMessage)GetMessage(ctx context.Context, id_message int) (*model.MessageService, error){
	msg, err := s.repoMessage.GetMessage(ctx, id_message)
	if err != nil{
		return nil, err
	}
	return msg, nil
}