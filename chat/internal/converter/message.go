package converter

import (
	"github.com/SigmarWater/messenger/chat/internal/model"
	pb "github.com/SigmarWater/messenger/chat/pkg/api/chat_service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MessageFromDesc(req *pb.SendMessageRequest) *model.MessageService{
	return &model.MessageService{
		ChatId: int(req.GetChatID()) ,
		FromUser: req.GetFrom(),
		TextMessage: req.GetText(),
	}
}

func MessageFromService(msg *model.MessageService) *pb.GetMessageResponse{
	return &pb.GetMessageResponse{
		ChatName: msg.ChatName,
		From: msg.FromUser, 
		Text: msg.TextMessage,
		Timestamp: timestamppb.New(msg.TimeAt),
	}
}