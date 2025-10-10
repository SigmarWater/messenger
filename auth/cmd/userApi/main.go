package main

import (
	"net"
	"github.com/SigmarWater/messenger/auth/internal/repository/users"
	pb "github.com/SigmarWater/messenger/auth/pkg/api/auth_service"
	"github.com/jackc/pgx/v4/pgxpool"

	"log"

	"context"
	"fmt"

	convFromClient "github.com/SigmarWater/messenger/auth/internal/converter"
	"github.com/SigmarWater/messenger/auth/internal/repository"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type User struct {
	name             string
	email            string
	password         string
	password_confirm string
	role             string
}

type UserApiServer struct {
	repUsers repository.UserRepository
	pb.UnimplementedUserAPIServer
}

const dbDNS string = "host=84.22.148.185 port=50000 user=sigmarwater password=sigmarwater dbname=users sslmode=disable"

func InitRepository(ctx context.Context, dns string) *pgxpool.Pool {
	pool, err := pgxpool.Connect(ctx, dns)
	if err != nil {
		log.Printf("Ошибка соединения с БД(users): %v\n", err)
		return nil
	}

	if err := pool.Ping(ctx); err != nil {
		log.Printf("Ошибка при проверке соединения с БД(users): %v\n", err)
		return nil
	}
	return pool
}

func NewUserApiServer() *UserApiServer {
	pool := InitRepository(context.Background(), dbDNS)
	if pool == nil {
		return nil
	}
	return &UserApiServer{repUsers: users.NewPostgresUserRepository(pool)}
}

func main() {
	server := grpc.NewServer()

	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Printf("Ошибка соединения: %v\n", err)
		return
	}

	pb.RegisterUserAPIServer(server, NewUserApiServer())

	reflection.Register(server)

	log.Println("Запускаем сервер")

	if err := server.Serve(lis); err != nil {
		log.Printf("Ошибка при прослушивании порта: %v\n", err)
		return
	}
}

func (u *UserApiServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	user := User{
		name:             req.GetName(),
		email:            req.GetEmail(),
		password:         req.GetPassword(),
		password_confirm: req.GetPasswordConfirm(),
		role:             req.GetRole().Enum().String(),
	}

	if user.password != user.password_confirm {
		errorStatus := status.New(codes.InvalidArgument, "Invalid password_confirm")
		ds, err := errorStatus.WithDetails(&errdetails.BadRequest{
			FieldViolations: []*errdetails.BadRequest_FieldViolation{
				{
					Field:       "Password",
					Description: fmt.Sprint("Password and Password Confirm aren't equal"),
				},
			},
		})

		if err != nil {
			return nil, errorStatus.Err()
		}

		return nil, ds.Err()
	}

	id, err := u.repUsers.InsertUser(ctx, convFromClient.ToUserFromDescCreate(req))
	
	if err != nil{
		log.Printf("Ошибка при добавлении user: %v\n", err)
		return nil, err
	}

	log.Printf("Created user with id: %+#v\n", id)

	return &pb.CreateResponse{Id: int64(id)}, nil
}

func (u *UserApiServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	user, err := u.repUsers.GetUser(ctx, int(req.GetId()))
	
	if err != nil{
		return nil, err
	}

	switch user.Role{
	case "user":
		 user.Role = "ROLE_USER"
	default:
		user.Role = "ROLE_ADMIN"
	}

	// if !ok {
	// 	errorStatus := status.New(codes.NotFound,
	// 		fmt.Sprintf("User with id: %d is not exists", id))

	// 	return nil, errorStatus.Err()
	// }

	role, ok := pb.Role_value[user.Role]

	if !ok {
		errorStatus := status.New(codes.Internal, "Such role isn't exists")
		return nil, errorStatus.Err()
	}

	log.Printf("Получаем user: %+#v\n", user)

	resp := &pb.GetResponse{
		Name:  user.Name,
		Email: user.Email,
		Role:  pb.Role(role),
		CreateAt: timestamppb.New(user.CreateAt),
		UpdateAt: timestamppb.New(user.UpdateAt),
	}

	return resp, nil
}

// func (u *UserApiServer) Update(ctx context.Context, req *pb.UpdateRequest) (*emptypb.Empty, error){
// 	id := req.GetId()

// 	user := User{
// 		name: req.GetName().Value,
// 		email: req.GetEmail().Value,
// 	}

// 	u.mutex.RLock()
// 	oldUser, ok := UserApi.users[int(id)]
// 	u.mutex.RUnlock()

// 	if !ok{
// 		errorStatus := status.New(codes.NotFound, codes.NotFound.String())
// 		return &emptypb.Empty{}, errorStatus.Err()
// 	}

// 	newUser := User{
// 		name: user.name,
// 		email: user.email,
// 		password: oldUser.password,
// 		password_confirm: oldUser.password_confirm,
// 		role: oldUser.role,
// 	}

// 	UserApi.mutex.Lock()

// 	UserApi.users[int(id)] = newUser

// 	UserApi.mutex.Unlock()

// 	log.Printf("Обновленные данные user:%d: %+#v\n", id, newUser)

// 	return &emptypb.Empty{}, nil
// }

// func (u *UserApiServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*emptypb.Empty, error){
// 	id := req.GetId()
// 	delete(UserApi.users, int(id))
// 	return &emptypb.Empty{}, nil
// }
