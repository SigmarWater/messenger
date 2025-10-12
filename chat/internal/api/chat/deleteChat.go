package chat

import(
	"context"
	pb "github.com/SigmarWater/messenger/chat/pkg/api/chat_service"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *pb.DeleteRequest) (*emptypb.Empty, error){
	err := i.chatSrv.DeleteChat(ctx, req.GetId())
	if err != nil{
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}
