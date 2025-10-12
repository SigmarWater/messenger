package messages

import (
	"context"
	"log"

	"github.com/SigmarWater/messenger/chat/internal/model"
	"github.com/SigmarWater/messenger/chat/internal/repository/messages/converter"
	rpModel "github.com/SigmarWater/messenger/chat/internal/repository/messages/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresMessageRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresMessageRepository(pool *pgxpool.Pool) *PostgresMessageRepository {
	return &PostgresMessageRepository{pool: pool}
}

func (p *PostgresMessageRepository) SendMessage(ctx context.Context, msg *model.MessageService) (int, error) {
	log.Printf("DEBUG: msg = %+v", msg)

	// Используем прямой SQL запрос вместо squirrel
	query := "INSERT INTO messages (id_chat, from_user, text_message, time_at) VALUES ($1, $2, $3, $4) RETURNING id_message"
	args := []interface{}{msg.ChatId, msg.FromUser, msg.TextMessage, msg.TimeAt}

	log.Printf("query: %v args: %v", query, args)

	var idMessage int
	err := p.pool.QueryRow(ctx, query, args...).Scan(&idMessage)
	if err != nil {
		log.Printf("Ошибка в запросе insert: %v\n", err)
		return 0, err
	}

	log.Printf("Получена запись с id: %d\n", idMessage)
	return idMessage, nil
}

func (p *PostgresMessageRepository) GetMessage(ctx context.Context, id_message int) (*model.MessageService, error) {
	log.Printf("DEBUG: id_message = %d (type: %T)", id_message, id_message)

	// Используем прямой SQL запрос вместо squirrel
	query := "SELECT id_message, id_chat, chat_name, from_user, text_message, time_at FROM messages INNER JOIN chats ON messages.id_chat = chats.id_chat WHERE id_message = $1"
	args := []interface{}{id_message}

	log.Printf("query: %v args: %v", query, args)

	var message rpModel.MessageRepository
	err := p.pool.QueryRow(ctx, query, args...).
		Scan(&message.IdMessage, &message.IdChat, &message.ChatName, &message.FromUser, &message.TextMessage, &message.TimeAt)
	if err != nil {
		log.Printf("Ошибка в запросе select: %v\n", err)
		return nil, err
	}

	return converter.ToMessageFromRepo(&message), nil
}
