package main

import (
	"net"

	rep "github.com/SigmarWater/messenger/auth/internal/repository/users"
	pb "github.com/SigmarWater/messenger/auth/pkg/api/auth_service"
	"github.com/jackc/pgx/v4/pgxpool"

	"log"

	"context"

	serv "github.com/SigmarWater/messenger/auth/internal/service/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	api "github.com/SigmarWater/messenger/auth/internal/api/auth"
)



const dbDNS string = "host=84.22.148.185 port=5430 user=sigmawater password=sigmawater dbname=messenger sslmode=disable"


func main() {
	server := grpc.NewServer()

	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Printf("Ошибка соединения: %v\n", err)
		return
	}

	pool, err := pgxpool.Connect(context.Background(), dbDNS)
	if err != nil{
		log.Printf("failed to connect to database: %v\n", err)
		return 
	}
	defer pool.Close()

	userRepo := rep.NewPostgresUserRepository(pool)
	userSrv := serv.NewService(userRepo)

	pb.RegisterUserAPIServer(server, api.NewImplementation(userSrv))

	reflection.Register(server)

	log.Println("Запускаем сервер")

	if err := server.Serve(lis); err != nil {
		log.Printf("Ошибка при прослушивании порта: %v\n", err)
		return
	}
}

