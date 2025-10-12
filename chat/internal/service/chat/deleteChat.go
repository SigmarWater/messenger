package chat

import(
	"context"
	"log"
)

func (s *servChat) DeleteChat(ctx context.Context, idChat int64) error{
	err := s.chatRepo.DeleteChat(ctx, idChat)
	if err != nil{
		log.Printf("Ошибка в service Delete Chat: %v\n", err)
	}
	return err
}