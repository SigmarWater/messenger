package main

import (
	"context"
	"log"
	"net"
	"github.com/SigmarWater/messenger/chat/internal/service"
	
	"github.com/SigmarWater/messenger/chat/internal/repository/chats"
	"github.com/SigmarWater/messenger/chat/internal/repository/messages"
	pb "github.com/SigmarWater/messenger/chat/pkg/api/chat_service"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	
	

	"github.com/SigmarWater/messenger/chat/internal/service/chat"
	"github.com/SigmarWater/messenger/chat/internal/service/message"
	api "github.com/SigmarWater/messenger/chat/internal/api/chat"
)


type Messenger struct{
	chatService service.ChatService
	messageService service.MessageService
	
	pb.UnimplementedChatApiServer	
}

const dbDNS string = "host=84.22.148.185 port=5430 user=sigmawater password=sigmawater dbname=messenger sslmode=disable"


func NewMessenger(chatService service.ChatService, messageService service.MessageService) *Messenger{
	return &Messenger{
		chatService: chatService,
		messageService: messageService,
	}
} 


func main(){
	server := grpc.NewServer()

	lis, err := net.Listen("tcp",":8085")
	if err != nil{
		log.Println("err")
		return 
	}

	pool, err := pgxpool.Connect(context.Background(), dbDNS)
	if err != nil{
		log.Fatal("failed to connect to database")
		return
	}
	defer pool.Close() 

	repoChat := chats.NewPostgresChatRepository(pool)
	repoMessage := messages.NewPostgresMessageRepository(pool)

	chatSrv := chat.NewServChat(repoChat)
	msgSrv := message.NewServMessage(repoMessage)

	pb.RegisterChatApiServer(server, api.NewImplementation(chatSrv, msgSrv))

	reflection.Register(server)

	log.Println("Server listen 8085")

	if err := server.Serve(lis); err != nil{
		log.Printf("Error serve %v\n", err)
		return 
	}
}