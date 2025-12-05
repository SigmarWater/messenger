package auth

import (
	"context"
	pb "github.com/SigmarWater/messenger/auth/pkg/api/auth_service"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"fmt"
	"github.com/SigmarWater/messenger/auth/internal/converter"
)

func (i *Implementation) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	user,err := converter.ToUserFromDescCreate(req)
	if err != nil{
		return nil, err
	}

	if user.EnterPassword != user.ConfirmPassword {
		errorStatus := status.New(codes.InvalidArgument, "Invalid password_confirm")
		ds, err := errorStatus.WithDetails(&errdetails.BadRequest{
			FieldViolations: []*errdetails.BadRequest_FieldViolation{
				{
					Field:       "Password",
					Description: fmt.Sprintln("Password and Password Confirm aren't equal"),
				},
			},
		})

		if err != nil {
			return nil, errorStatus.Err()
		}

		return nil, ds.Err()
	}

	id, err := i.userService.Create(ctx, user)

	if err != nil {
		log.Printf("Ошибка при добавлении user: %v\n", err)
		return &pb.CreateResponse{Id:id}, err
	}

	log.Printf("Created user with id: %+#v\n", id)

	return &pb.CreateResponse{Id: int64(id)}, nil
}