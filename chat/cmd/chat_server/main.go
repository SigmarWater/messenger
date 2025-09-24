package main

import (
	"context"
	"io"
	"log"
	"net"
	"sync"
	"fmt"

	pb "github.com/SigmarWater/messenger/chat/pkg/api/chat_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/grpc/status"
)

type Message struct{
	from string 
	text string 
	timestamp *timestamppb.Timestamp 
}

type Chat struct{
	usernames []string 
	messages []Message 
}

type Messenger struct{
	chats map[int]*Chat
	mutex sync.RWMutex
	pb.UnimplementedChatApiServer	
}

func NewMessenger() *Messenger{
	return &Messenger{
		chats: make(map[int]*Chat),
	}
} 

func NewChat(usernames []string) *Chat{
	return &Chat{
		usernames: usernames,
		messages: make([]Message, 0),
	}
}

var messenger = NewMessenger()

func main(){
	server := grpc.NewServer()

	lis, err := net.Listen("tcp",":8085")
	if err != nil{
		log.Println("err")
		return 
	}

	pb.RegisterChatApiServer(server, messenger)

	reflection.Register(server)

	log.Println("Server listen 8085")

	if err := server.Serve(lis); err != nil{
		log.Printf("Error serve %v\n", err)
		return 
	}
}

func (m *Messenger) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error){
	usernames := req.GetUsernames()

	chat := NewChat(usernames)

	id := len(m.chats) + 1


	m.mutex.Lock()
	m.chats[id] = chat
	m.mutex.Unlock()
	

	log.Printf("Успешно создан чат c id: %d\n", id)

	return &pb.CreateResponse{Id: int64(id)}, nil
}


func (m *Messenger) Delete(ctx context.Context, req *pb.DeleteRequest) (*emptypb.Empty, error){
	id := req.GetId()

	if _, ok := m.chats[int(id)]; !ok{
		errorStatus:= status.New(codes.NotFound, 
			fmt.Sprintf("Чата с id: %d - не существует\n", id))

		return &emptypb.Empty{}, errorStatus.Err()
	}

	delete(m.chats, int(id)) 

	log.Printf("Успешно удален час с id: %d\n", id)

	return &emptypb.Empty{}, nil
}

func (m *Messenger) SendMessage(stream grpc.ClientStreamingServer[pb.SendMessageRequest, wrapperspb.StringValue]) error{
	sendMessages := 0
	for{
		message, err := stream.Recv() 

		if err == io.EOF{
			log.Println("Все сообщения прочитаны")
			return stream.SendAndClose(&wrapperspb.StringValue{
				Value: fmt.Sprintf("Отправлено %d сообщений", sendMessages),})
		}

		if err != nil{
			errStatus := status.New(codes.Aborted, fmt.Sprintf("Error in reading: %v", err))
			return errStatus.Err() 
		}

		

		chatID := message.GetChatID() 
		
		if _, ok := m.chats[int(chatID)]; !ok{
			log.Printf("Not found chat with id: %d\n", chatID)
			errStatus := status.New(codes.NotFound, fmt.Sprintf("Not found chat with id: %d", chatID))
			return errStatus.Err()
		}

		chat := m.chats[int(chatID)]

		messageChat := Message{
			from: message.GetFrom(),
			text: message.GetText(),
			timestamp: message.GetTimestamp(),
		}

		chat.messages = append(chat.messages, messageChat)
		log.Printf("В чат %d прилетело сообщение от %s\n", chatID, message.GetFrom())
		sendMessages++
	}
}