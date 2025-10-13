package chats

import (
	"context"
	"log"

	"github.com/SigmarWater/messenger/chat/internal/model"
	"github.com/SigmarWater/messenger/chat/internal/repository/chats/converter"
	repoModel "github.com/SigmarWater/messenger/chat/internal/repository/chats/model"
	"github.com/jackc/pgx/v4/pgxpool"
	sq "github.com/Masterminds/squirrel"
)

type PostgresChatRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresChatRepository(pool *pgxpool.Pool) *PostgresChatRepository {
	return &PostgresChatRepository{pool: pool}
}

// DeleteChat(ctx context.Context, idChat int)

func (p *PostgresChatRepository) CreateChat(ctx context.Context, chatInfo *model.ChatService) (*model.ChatService, error) {
	builderInsert := sq.Insert("chats").
	PlaceholderFormat(sq.Dollar).
	Columns("chat_name").
	Values(chatInfo.ChatName).
	Suffix("RETURNING id_chat, chat_name")

	query, args, err := builderInsert.ToSql()
	if err != nil{
		log.Printf("Ошибка в преобразовании запроса CreateChat: %v\n", err)
		return nil, err
	}

	var chatRep repoModel.ChatRepository
	err = p.pool.QueryRow(ctx, query, args...).Scan(&chatRep.IdChat, &chatRep.ChatName)
	if err != nil {
		log.Printf("Ошибка выполнения запроса в CreateChat: %v\n", err)
		return nil, err
	}

	return converter.MsgRepoToService(chatRep), nil
}

func (p *PostgresChatRepository) DeleteChat(ctx context.Context, idChat int64) error {

	duilderDelete := sq.Delete("chats").
	PlaceholderFormat(sq.Dollar).
	Where(sq.Eq{"id_chat": idChat})

	query, args, err := duilderDelete.ToSql()
	if err != nil{
		log.Printf("Ошибка при создании запроса: %v\n", err)
		return err 
	}
	
	_, err = p.pool.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("Ошибка выполнения запроса в DeleteChat: %v\n", err)
		return err
	}

	return nil
}
