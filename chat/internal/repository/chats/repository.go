package chats

import (
	"context"
	"log"

	"github.com/SigmarWater/messenger/chat/internal/model"
	"github.com/SigmarWater/messenger/chat/internal/repository/chats/converter"
	repoModel "github.com/SigmarWater/messenger/chat/internal/repository/chats/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresChatRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresChatRepository(pool *pgxpool.Pool) *PostgresChatRepository {
	return &PostgresChatRepository{pool: pool}
}

// DeleteChat(ctx context.Context, idChat int)

func (p *PostgresChatRepository) CreateChat(ctx context.Context, chatInfo *model.ChatService) (*model.ChatService, error) {
	log.Printf("DEBUG: chatInfo.ChatName = '%s' (type: %T)", chatInfo.ChatName, chatInfo.ChatName)

	// Используем прямой SQL запрос вместо squirrel
	query := "INSERT INTO chats (chat_name) VALUES ($1) RETURNING id_chat, chat_name"
	args := []interface{}{chatInfo.ChatName}

	log.Printf("query: %v args: %v", query, args)

	var chatRep repoModel.ChatRepository
	err := p.pool.QueryRow(ctx, query, args...).Scan(&chatRep.IdChat, &chatRep.ChatName)
	if err != nil {
		log.Printf("Ошибка выполнения запроса в CreateChat: %v\n", err)
		return nil, err
	}

	return converter.MsgRepoToService(chatRep), nil
}

func (p *PostgresChatRepository) DeleteChat(ctx context.Context, idChat int64) error {
	log.Printf("DEBUG: idChat = %d (type: %T)", idChat, idChat)

	// Используем прямой SQL запрос вместо squirrel
	query := "DELETE FROM chats WHERE id_chat = $1"
	args := []interface{}{idChat}

	log.Printf("query: %v args: %v", query, args)

	_, err := p.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("Ошибка выполнения запроса в DeleteChat: %v\n", err)
		return err
	}

	return nil
}
