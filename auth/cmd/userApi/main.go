package main

import (
	"net"

	pb "github.com/SigmarWater/messenger/auth/pkg/api/auth_service"

	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main(){
	server := grpc.NewServer()

	lis, err := net.Listen("tcp", ":8080")
	if err != nil{
		log.Printf("Ошибка соединения: %v\n", err)
		return 
	}
	
	pb.RegisterUserAPIServer(server, pb.UnimplementedUserAPIServer{})
	
	reflection.Register(server)

	log.Println("Запускаем сервер")

	if err := server.Serve(lis); err != nil{
		log.Printf("Ошибка при прослушивании порта: %v\n", err)
		return
	}
}