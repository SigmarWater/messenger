package main

import (
	pb "github.com/SigmarWater/messenger/jwt/pkg/api/jwt_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"log"
	"github.com/SigmarWater/messenger/jwt/internal/api"
)

const(
	grpcPort = ":10000"
)

func main() {
	srv	 := grpc.NewServer()

	pb.RegisterJWTServiceServer(srv, api.NewJWTHandler())

	reflection.Register(srv)

	listener, err := net.Listen("tcp", grpcPort)
	if err != nil{
		log.Fatal("error listen:", err.Error())
	}

	log.Printf("Start server on port: %s\n", grpcPort)

	if err := srv.Serve(listener); err != nil{
		log.Printf("Failed to serve grpc Server: %v", err.Error())
	}
}