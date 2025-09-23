package main

import (
	"net"
	"sync"

	pb "github.com/SigmarWater/messenger/auth/pkg/api/auth_service"

	"log"

	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type User struct{
	name string 
	email string 
	password string 
	password_confirm string 
	role string 
}

type UserApiServer struct{
	id int
	users map[int]User 
	mutex sync.RWMutex
	pb.UnimplementedUserAPIServer
}

func NewUserApiServer () *UserApiServer{
	return &UserApiServer{users: make(map[int]User)}
}

var UserApi = NewUserApiServer()

func main(){
	server := grpc.NewServer()

	lis, err := net.Listen("tcp", ":8083")
	if err != nil{
		log.Printf("Ошибка соединения: %v\n", err)
		return 
	}

	pb.RegisterUserAPIServer(server, UserApi)
	
	reflection.Register(server)

	log.Println("Запускаем сервер")

	if err := server.Serve(lis); err != nil{
		log.Printf("Ошибка при прослушивании порта: %v\n", err)
		return
	}
}

func (u *UserApiServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error){
	
	user := User{
		name: req.GetName(), 
		email: req.GetEmail(),
		password: req.GetPassword(),
		password_confirm: req.GetPasswordConfirm(),
		role: req.GetRole().Enum().String(),
	}
	
	UserApi.mutex.Lock()
	UserApi.users[UserApi.id + 1] = user
	UserApi.mutex.Unlock()

	UserApi.id++
	
	log.Printf("Created user: %+#v\n", user)

	return &pb.CreateResponse{Id:int64(UserApi.id)}, nil
}

func (u *UserApiServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error){
	id := req.GetId()

	UserApi.mutex.RLock()
	user, ok := UserApi.users[int(id)]
	UserApi.mutex.RUnlock()

	if !ok{
		return nil, errors.New("Now found")
	}

	role, ok := pb.Role_value[user.role]
	if !ok{
		return nil, errors.New("Нет такого значения роли")
	}

	log.Printf("Получаем user: %+#v\n", user)

	resp := &pb.GetResponse{
		Id: id, 
		Name: user.name,
		Email: user.email,
		Role: pb.Role(role),
	}

	return resp, nil  
}

func (u *UserApiServer) Update(ctx context.Context, req *pb.UpdateRequest) (*emptypb.Empty, error){
	id := req.GetId()

	user := User{
		name: req.GetName().Value,
		email: req.GetEmail().Value,
	}


	u.mutex.RLock()
	oldUser := UserApi.users[int(id)]
	u.mutex.RUnlock()

	newUser := User{
		name: user.name, 
		email: user.email, 
		password: oldUser.password,
		password_confirm: oldUser.password_confirm,
		role: oldUser.role,
	}

	UserApi.mutex.Lock()

	UserApi.users[int(id)] = newUser

	UserApi.mutex.Unlock()

	log.Printf("Обновленные данные user:%d: %+#v\n", id, user)

	return &emptypb.Empty{}, nil
}

func (u *UserApiServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*emptypb.Empty, error){
	id := req.GetId()
	delete(UserApi.users, int(id)) 
	return &emptypb.Empty{}, nil
}