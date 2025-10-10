package converter

import (
	"time"

	"github.com/SigmarWater/messenger/auth/internal/model"
	pb "github.com/SigmarWater/messenger/auth/pkg/api/auth_service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func roleToString(role pb.Role) string {
	switch role {
	case pb.Role_ROLE_USER:
		return "user"
	case pb.Role_ROLE_ADMIN:
		return "admin"
	default:
		return "unspecified"
	}
}

func stringToRole(role string) pb.Role {
	switch role {
	case "user":
		return pb.Role_ROLE_USER
	case "admin":
		return pb.Role_ROLE_ADMIN
	default:
		return pb.Role_ROLE_UNSPECIFIED
	}
}

func ToUserFromDescCreate(req *pb.CreateRequest) *model.UserService {
	return &model.UserService{
		Name:          req.GetName(),
		Email:         req.GetEmail(),
		EnterPassword: req.GetPassword(),
		Role:          roleToString(req.GetRole()),
		CreateAt:      time.Now(),
	}
}

func ToDescFromUser(user model.UserService) *pb.GetResponse {
	return &pb.GetResponse{
		Name:     user.Name,
		Email:    user.Email,
		Role:     stringToRole(user.Role),
		CreateAt: timestamppb.New(user.CreateAt),
		UpdateAt: timestamppb.New(user.UpdateAt),
	}
}
