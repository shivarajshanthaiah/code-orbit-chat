package service

import (
	"context"
	"log"
	"time"

	"github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/models"
	pb "github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/proto"
	"github.com/shivaraj-shanthaiah/code_orbit_chat/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ch *ChatService) AddCommentService(ctx context.Context, req *pb.CommentRequest) (*pb.CommentResponse, error) {
	comment := models.Comment{
		ID:              primitive.NewObjectID().Hex(),
		ProblemID:       int(req.ProblemId),
		UserID:          req.UserId,
		Content:         req.Content,
		ParentCommentID: "",
		Timestamp:       time.Now(),
	}

	err := ch.Repo.AddComment(ctx, &comment)
	if err != nil {
		return nil, err
	}

	return &pb.CommentResponse{
		Message: "Comment added successfully",
		Comment: &pb.Comment{
			Id:              comment.ID,
			ProblemId:       uint32(comment.ProblemID),
			UserId:          comment.UserID,
			Content:         comment.Content,
			ParentCommentId: comment.ParentCommentID,
			Timestamp:       comment.Timestamp.Format(time.RFC3339),
		},
	}, nil
}

func (ch *ChatService) ReplyToCommentService(ctx context.Context, req *pb.ReplyRequest) (*pb.CommentResponse, error) {
	reply := models.Comment{
		ID:              primitive.NewObjectID().Hex(),
		UserID:          req.UserId,
		Content:         req.Content,
		ParentCommentID: req.CommentId,
		Timestamp:       time.Now(),
	}

	err := ch.Repo.AddReply(ctx, req.CommentId, &reply)
	if err != nil {
		return nil, err
	}

	return &pb.CommentResponse{
		Message: "Reply added successfully",
		Comment: &pb.Comment{
			Id:              reply.ID,
			UserId:          reply.UserID,
			Content:         reply.Content,
			ParentCommentId: reply.ParentCommentID,
			Timestamp: reply.Timestamp.Format(time.RFC3339),
		},
	}, nil
}

func (ch *ChatService) GetCommentsForProblemService(ctx context.Context, req *pb.FetchCommentsRequest) (*pb.FetchCommentsResponse, error) {
	comments, err := ch.Repo.GetCommentsByProblemID(ctx, int(req.ProblemId))
	if err != nil {
		return nil, err
	}

	var grpcComments []*pb.Comment
	for _, comment := range comments {
		grpcComments = append(grpcComments, utils.ConvertToGRPCCommentForProblem(comment))
	}

	return &pb.FetchCommentsResponse{
		Comments: grpcComments,
	}, nil
}

func (ch *ChatService) GetUserCommentsService(ctx context.Context, req *pb.FetchUserCommentsRequest) (*pb.FetchUserCommentsResponse, error) {
	if req.UserId == "" {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	comments, err := ch.Repo.GetUserComments(ctx, req.UserId)
	if err != nil {
		log.Printf("Error getting comments: %v", err)
		return nil, status.Error(codes.Internal, "failed to fetch comments")
	}

	// Convert the comments to the gRPC format
	grpcComments := make([]*pb.Comment, 0, len(comments))
	for _, comment := range comments {
		grpcComments = append(grpcComments, utils.ConvertToGRPCCommentForComment(comment))
	}

	return &pb.FetchUserCommentsResponse{
		Comments: grpcComments,
	}, nil
}
