package main

import (
	"net"
	"os"

	rep "github.com/SigmarWater/messenger/auth/internal/repository/users"
	pb "github.com/SigmarWater/messenger/auth/pkg/api/auth_service"
	"github.com/jackc/pgx/v4/pgxpool"

	"log"

	"context"

	serv "github.com/SigmarWater/messenger/auth/internal/service/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"flag"

	api "github.com/SigmarWater/messenger/auth/internal/api/auth"
	"github.com/SigmarWater/messenger/auth/internal/config"
	"github.com/SigmarWater/messenger/auth/internal/config/env"
)

var serviceConf string

func init() {
	flag.StringVar(&serviceConf, "env", ".env", "path to config")
}

func main() {
	flag.Parse()

	err := config.Load(serviceConf)
	if err != nil {
		log.Fatalf("failed download env: %v\n", err)
	}

	server := grpc.NewServer()
	grpcConfig, err := env.NewGrpcConfig()
	if err != nil {
		log.Fatalf("bad config for grpc: %v\n", err)
	}

	pgCongig, err := env.NewPgConfig()
	if err != nil {
		log.Fatalf("bad config for postgres: %v\n", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Printf("Ошибка соединения: %v\n", err)
		return
	}

	pool, err := pgxpool.Connect(context.Background(), pgCongig.DNS())
	if err != nil {
		log.Printf("failed to connect to database: %v\n", err)
		return
	}
	defer pool.Close()

	userRepo := rep.NewPostgresUserRepository(pool)
	userSrv := serv.NewService(userRepo)

	pb.RegisterUserAPIServer(server, api.NewImplementation(userSrv))

	reflection.Register(server)

	log.Printf("Запускаем сервер по адресу: %v\n", net.JoinHostPort(os.Getenv("GRPC_HOST"), os.Getenv("GRPC_PORT")))

	if err := server.Serve(lis); err != nil {
		log.Printf("Ошибка при прослушивании порта: %v\n", err)
		return
	}
}
