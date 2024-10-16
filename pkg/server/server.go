package server

import (
	"fmt"
	"log"
	"net"

	"github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/handler"
	pb "github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/proto"
	"google.golang.org/grpc"
)

func NewGrpcChatServer(port string, handlr *handler.ChatServiceServer) error {
	log.Println("connecting to gRPC server")
	addr := fmt.Sprintf(":%s", port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("error creating listener on %v", addr)
		return err
	}

	grpc := grpc.NewServer()

	pb.RegisterChatServiceServer(grpc, handlr)
	log.Printf("listening on gRPC server %v", port)
	err = grpc.Serve(lis)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}
	return nil
}
