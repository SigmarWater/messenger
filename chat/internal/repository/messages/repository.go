package message 

import(
	"github.com/SigmarWater/messenger/chat/internal/model"
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	rpModel "github.com/SigmarWater/messenger/chat/internal/repository/messages/model"
)

const dbDNS = "host=84.22.148.185 port=50000 user=sigmawater password=sigmawater dbname=messenger sslmode=disable"

func SendMessage (ctx context.Context, msg *model.MessageService) (int, error){
	pool, err :=  pgxpool.Connect(ctx, dbDNS)
	if err != nil{
		log.Printf("Ошибка соединение: %v\n", err)
		return 0, err 
	}

	if err := pool.Ping(ctx); err != nil{
		log.Printf("Ошибка подключения: %v\n", err)
		return 0, nil 
	}

	builderInsert := sq.Insert("messages").
	PlaceholderFormat(sq.Dollar).
	Columns("id_chat", "from_user", "text_message", "time_at").
	Values(msg.ChatId, msg.From_user, msg.Text_message, msg.Time_at).
	Suffix("RETURNIG id_message")

	query, arguments, err := builderInsert.ToSql()
	if err != nil{
		log.Printf("Ошибка при создании запроса insert: %v\n", err)
		return 0, err
	}

	var idMessage int
	err = pool.QueryRow(ctx, query, arguments...).Scan(&idMessage)
	if err != nil{
		log.Printf("Ошибка в запросе insert: %v\n", err)
		return 0, err
	}

	log.Printf("Получена запись с id: %d\n", idMessage)
	return idMessage, nil
}

func GetMessage (ctx context.Context, id_message int) (*model.MessageService, error){
	pool, err := pgxpool.Connect(ctx, dbDNS)
	if err != nil{
		log.Printf("Ошибка соединение: %v\n", err)
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil{
		log.Printf("Ошибка подключения: %v\n", err)
		return nil, err
	}

	
	builderSelect := sq.Select("id_chat", "from_user", "text_message", "time_at").
	From("messages").
	PlaceholderFormat(sq.Dollar).
	Where(sq.Eq{"id_message":id_message})

	query, args, err := builderSelect.ToSql()

	if err != nil{
		log.Printf("Ошибка при создании запроса select: %v\n", err)
		return nil, err
	}

	var message *rpModel.MessageRepository
	err = pool.QueryRow(ctx, query, args).Scan(message.Id_chat, message.From_user, message.Text_message, message.Time_at)
	if err != nil{
		log.Printf("Ошибка в запросе select: %v\n", err)
		return nil, err
	}

	return 

}