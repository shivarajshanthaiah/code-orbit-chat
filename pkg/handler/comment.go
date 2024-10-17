package handler

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/proto"
)

func (ch *ChatServiceServer) AddComment(ctx context.Context, p *pb.CommentRequest) (*pb.CommentResponse, error) {
	response, err := ch.svc.AddCommentService(ctx, p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (ch *ChatServiceServer) ReplyToComment(ctx context.Context, p *pb.ReplyRequest) (*pb.CommentResponse, error) {
	response, err := ch.svc.ReplyToCommentService(ctx, p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (ch *ChatServiceServer) FetchComments(ctx context.Context, p *pb.FetchCommentsRequest) (*pb.FetchCommentsResponse, error) {
	response, err := ch.svc.GetCommentsForProblemService(ctx, p)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (ch *ChatServiceServer) FetchUserComments(ctx context.Context, p *pb.FetchUserCommentsRequest) (*pb.FetchUserCommentsResponse, error){
	response, err := ch.svc.GetUserCommentsService(ctx, p)
	if err != nil{
		return response, err
	}
	return response, nil
}