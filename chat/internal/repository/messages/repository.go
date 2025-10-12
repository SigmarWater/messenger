package messages

import(
	"github.com/SigmarWater/messenger/chat/internal/model"
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	rpModel "github.com/SigmarWater/messenger/chat/internal/repository/messages/model"
	"github.com/SigmarWater/messenger/chat/internal/repository/messages/converter"
)

type PostgresMessageRepository struct{
	pool *pgxpool.Pool
}

func NewPostgresMessageRepository(pool *pgxpool.Pool) *PostgresMessageRepository{
	return &PostgresMessageRepository{pool:pool}
}

func (p *PostgresMessageRepository)SendMessage (ctx context.Context, msg *model.MessageService) (int, error){
	builderInsert := sq.Insert("messages").
	PlaceholderFormat(sq.Dollar).
	Columns("id_chat", "from_user", "text_message", "time_at").
	Values(msg.ChatId, msg.FromUser, msg.TextMessage, msg.TimeAt).
	Suffix("RETURNIG id_message")

	query, arguments, err := builderInsert.ToSql()
	if err != nil{
		log.Printf("Ошибка при создании запроса insert: %v\n", err)
		return 0, err
	}

	var idMessage int
	err = p.pool.QueryRow(ctx, query, arguments...).Scan(&idMessage)
	if err != nil{
		log.Printf("Ошибка в запросе insert: %v\n", err)
		return 0, err
	}

	log.Printf("Получена запись с id: %d\n", idMessage)
	return idMessage, nil
}

func (p *PostgresMessageRepository)GetMessage (ctx context.Context, id_message int) (*model.MessageService, error){
	builderSelect := sq.Select("id_message","id_chat", "chat_name","from_user", "text_message", "time_at").
	From("messages").
	PlaceholderFormat(sq.Dollar).
	Where(sq.Eq{"id_message":id_message}).
	InnerJoin("chats", "messages.id_chat", "chats.id_chat")

	query, args, err := builderSelect.ToSql()

	if err != nil{
		log.Printf("Ошибка при создании запроса select: %v\n", err)
		return nil, err
	}

	var message rpModel.MessageRepository
	err = p.pool.QueryRow(ctx, query, args).
	Scan(&message.IdMessage, &message.IdChat, &message.ChatName, &message.FromUser, &message.TextMessage, &message.TimeAt)
	if err != nil{
		log.Printf("Ошибка в запросе select: %v\n", err)
		return nil, err
	}

	return converter.ToMessageFromRepo(&message), nil
}