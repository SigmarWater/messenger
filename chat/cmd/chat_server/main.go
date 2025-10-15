package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/SigmarWater/messenger/chat/internal/service"

	"github.com/SigmarWater/messenger/chat/internal/repository/chats"
	"github.com/SigmarWater/messenger/chat/internal/repository/messages"
	pb "github.com/SigmarWater/messenger/chat/pkg/api/chat_service"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	api "github.com/SigmarWater/messenger/chat/internal/api/chat"
	"github.com/SigmarWater/messenger/chat/internal/service/chat"
	"github.com/SigmarWater/messenger/chat/internal/service/message"
	"github.com/SigmarWater/messenger/chat/internal/config"
	"github.com/SigmarWater/messenger/chat/internal/config/env"
)

var configPath string

func init(){
	flag.StringVar(&configPath, "env", "c:\\users\\admin\\desktop\\goproject\\messenger\\postgres\\migrations\\.env", "path to config file")
}

type Messenger struct{
	chatService service.ChatService
	messageService service.MessageService
	
	pb.UnimplementedChatApiServer	
}

func NewMessenger(chatService service.ChatService, messageService service.MessageService) *Messenger{
	return &Messenger{
		chatService: chatService,
		messageService: messageService,
	}
} 


func main(){

	flag.Parse()

	err := config.Load(configPath)
	if err != nil{
		log.Fatalf("failed to load config: %v\n", err)
	}

	grpcConfig, err := env.NewConfig()
	if err != nil{
		log.Fatalf("failed to get grpc config %v\n", err)
	}

	pgConfig, err := env.NewPGConfig() 
	if err != nil{
		log.Fatalf("failed to listen %v\n", err)
	}

	server := grpc.NewServer()

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil{
		log.Println("err")
		return 
	}

	pool, err := pgxpool.Connect(context.Background(), pgConfig.DSN())
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