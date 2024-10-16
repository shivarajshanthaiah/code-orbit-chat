package handler

import (
	pb "github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/proto"
	"github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/service/interfaces"
)

type ChatServiceServer struct {
	pb.UnimplementedChatServiceServer
	svc interfaces.ChatServiceInter
}

func NewChatServiceServer(svc interfaces.ChatServiceInter) *ChatServiceServer {
	return &ChatServiceServer{
		svc: svc,
	}
}
