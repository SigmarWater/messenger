package chats 

import(
	"github.com/SigmarWater/messenger/chat/internal/model"
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	rpModel "github.com/SigmarWater/messenger/chat/internal/repository/chats/model"
)

const dbDNS = "host=84.22.148.185 port=50000 user=sigmawater password=sigmawater dbname=messenger sslmode=disable"

func GetInfoChat(ctx context.Context, id_chat int) (*model.ChatService, error){
	pool, err := pgxpool.Connect(ctx, dbDNS)
	if err != nil{
		log.Printf("Ошибка соединение: %v\n", err)
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil{
		log.Printf("Ошибка подключения: %v\n", err)
		return nil, err
	}

	
	builderSelect := sq.Select("chat_name").
	From("chats").
	PlaceholderFormat(sq.Dollar).
	Where(sq.Eq{"id_message":id_chat})

	query, args, err := builderSelect.ToSql()

	if err != nil{
		log.Printf("Ошибка при создании запроса select: %v\n", err)
		return nil, err
	}

	var chat *rpModel.ChatRepository
	err = pool.QueryRow(ctx, query, args).Scan(chat.ChatName)
	if err != nil{
		log.Printf("Ошибка в запросе select: %v\n", err)
		return nil, err
	}

	return 
}