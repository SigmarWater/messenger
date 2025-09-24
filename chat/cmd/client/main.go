package main

import (
	"context"
	"log"

	pb "github.com/SigmarWater/messenger/chat/pkg/api/chat_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)


func main() {
	clientConn, err := grpc.NewClient(":8085", 
	grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil{
		log.Println("Error connect client")
		return 
	}

	client := pb.NewChatApiClient(clientConn)

	res, err := client.Create(context.Background(), &pb.CreateRequest{
		Usernames: []string{"Artem Udalcov", "Kirill Egashin", "Alexandr Ozerskiy"},
	})

	if err != nil{
		log.Printf("Ошибка: %v\n", err)
		return 
	}

	log.Printf("Создан чат c id: %d\n", res.GetId())

	id := 3
	_, err = client.Delete(context.Background(), &pb.DeleteRequest{Id: int64(id)})
	if err != nil{
		switch status.Code(err){
		case codes.NotFound:
			log.Printf("Не существует чата с id:%d\n", id)
			return
		default:
			log.Printf("Error: %v\n", err)
			return
		}
	}
	log.Printf("Чат с id:%d успешно удален\n", id)



	stream, err := client.SendMessage(context.Background())
	if err != nil{
		log.Println("error", err.Error())
		return
	}

	message1 := &pb.SendMessageRequest{
		ChatID: 1,
		From: "Artem Udalcov",
		Text: "Долбаеб привет, как твои дела нахуй?",
		Timestamp: timestamppb.Now(),
	}

	message2 := &pb.SendMessageRequest{
		ChatID: 1,
		From: "Кirill Egishin",
		Text: "Сам долбаеб",
		Timestamp: timestamppb.Now(),
	}

	err = stream.Send(message1) 
	if err != nil{
		log.Println(err)
		return 
	}
	

	err = stream.Send(message2)
	if err != nil{
		log.Println(err)
		return 
	}
	

	resStr, err := stream.CloseAndRecv()
	if err != nil{
		log.Println("Ошибка после close")
		return 
	}

	log.Println(resStr)
}