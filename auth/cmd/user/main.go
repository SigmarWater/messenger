package main

import (
	"context"
	"log"
	"google.golang.org/grpc/status"

	pb "github.com/SigmarWater/messenger/auth/pkg/api/auth_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"google.golang.org/genproto/googleapis/rpc/errdetails"

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
		switch status.Code(err){
		case codes.InvalidArgument:
			log.Println("Некорректные данные")
		default:
			log.Printf("Ошибка %v", err)
		}

		if st, ok := status.FromError(err); ok{
			log.Printf("code: %v message: %v\n", st.Code(), st.Message())
			for _, d := range st.Details(){
				switch detail := d.(type){
				case *errdetails.BadRequest_FieldViolation:
					log.Printf("details: %v\n", detail)
				default:
					log.Printf("detais: %v\n", detail)
				}
			}
		}else{
			log.Println("not grpc")
		}

		return
	}
	
	log.Printf("Res : %v\n", resCreate)

	resGet, err := client.Get(context.Background(), &pb.GetRequest{Id: 1})

	if err != nil{
		switch status.Code(err){
		case codes.NotFound:
			log.Println("Not found user") 
		case codes.Internal:
			log.Println("Ошибка внутри сервера")
		default:
			log.Printf("Error: %v\n", err)
		}

		if st, ok := status.FromError(err); ok{
			log.Printf("code: %v, message: %v\n", 
			st.Code(), st.Message())
			for _, d := range st.Details(){
				switch d.(type){
				case *errdetails.BadRequest_FieldViolation:
					log.Printf("details: %v\n", d)
				default:
					log.Printf("details: %v\n", d)
				}
			}
		}else{
			log.Println("not grpc")
		}

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
		switch status.Code(err){
		case codes.NotFound:
			log.Println("Not found user")
		default:
			log.Printf("Error: %v\n", err)
		}

		if st, ok := status.FromError(err); ok{
			log.Printf("code: %v, message: %v\n", 
			st.Code(), st.Message())

			for _, d := range st.Details(){
				switch d.(type){
				case *errdetails.BadRequest_FieldViolation:
					log.Printf("details: %v\n", d)
				default:
					log.Printf("details: %v\n", d)
				}
			}
		}else{
			log.Println("not grpc")
		}

		return 
	}

	log.Printf("Res Update: %v\n", resUpdate)

	resGet2, err := client.Get(context.Background(), &pb.GetRequest{Id: 1})

	if err != nil{
		switch status.Code(err){
		case codes.NotFound:
			log.Println("Not found user") 
		case codes.Internal:
			log.Println("Ошибка внутри сервера")
		default:
			log.Printf("Error: %v\n", err)
		}

		if st, ok := status.FromError(err); ok{
			log.Printf("code: %v, message: %v\n", 
			st.Code(), st.Message())

			for _, d := range st.Details(){
				switch d.(type){
				case *errdetails.BadRequest_FieldViolation:
					log.Printf("details: %v\n", d)
				default:
					log.Printf("details: %v\n", d)
				}
			}
		}else{
			log.Println("not grpc")
		}
		return
	}

	log.Printf("Res Get2: %v\n", resGet2)

}