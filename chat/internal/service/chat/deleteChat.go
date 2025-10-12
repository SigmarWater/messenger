package chat

import(
	"context"
)

func (s *servChat) DeleteChat(ctx context.Context, idChat int) error{
	return s.chatRepo.DeleteChat(ctx, idChat)
}