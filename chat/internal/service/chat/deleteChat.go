package chat

import(
	"context"
)

func (s *servChat) DeleteChat(ctx context.Context, idChat int64) error{
	return s.chatRepo.DeleteChat(ctx, idChat)
}