package message

import(
	"context"
	"github.com/SigmarWater/messenger/chat/internal/model"
)

func (s *servMessage)SendMessage(ctx context.Context, msg *model.MessageService) (id int, err error){
	id, err = s.repoMessage.SendMessage(ctx, msg)
	if err != nil{
		return 0, err
	}
	
	return id, nil
}