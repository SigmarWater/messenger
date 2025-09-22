package main

import (
	pb "github.com/SigmarWater/messenger/auth/pkg/api/auth_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserAPI struct{
	pb.UnimplementedUserAPIServer 
}

func main(){
	conn, err := grpc.NewClient(":8080", 
	grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		return 
	}
}