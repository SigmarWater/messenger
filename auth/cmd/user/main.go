package main

import (
	"context"
	"log"

	pb "github.com/SigmarWater/messenger/auth/pkg/api/auth_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)


func main(){
	conn, err := grpc.NewClient(":8083", 
	grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Println("Error create client")
		return 
	}

	client := pb.NewUserAPIClient(conn)

	user := pb.CreateRequest{
		Name: "Artem",
		Email: "artemudalcov05@gmail.com",
		Password: "qwerty",
		PasswordConfirm: "qwerty",
		Role: pb.Role_ROLE_ADMIN,
	}

	resCreate, err := client.Create(context.Background(), &user)
	if err != nil{
		log.Printf("Error of creating client: %v\n", err.Error())
		return 
	}
	log.Printf("Res : %v\n", resCreate)

	resGet, err := client.Get(context.Background(), &pb.GetRequest{Id: 1})
	if err != nil{
		log.Printf("Error of getting client: %v\n", err.Error())
		return
	}

	log.Printf("Res Get: %v\n", resGet)

	newUser := pb.UpdateRequest{
		Id: 1, 
		Name: &wrapperspb.StringValue{Value:"SigmarWater"},
		Email: &wrapperspb.StringValue{Value:"sigmarwaterofficial.com"},
	}
	resUpdate, err := client.Update(context.Background(), &newUser)
	if err != nil{
		log.Printf("Error updating user: %v\n", err)
		return 
	}
	log.Printf("Res Update: %v\n", resUpdate)

	resGet2, err := client.Get(context.Background(), &pb.GetRequest{Id: 1})
	if err != nil{
		log.Printf("Error of getting client: %v\n", err.Error())
		return
	}

	log.Printf("Res Get2: %v\n", resGet2)

}