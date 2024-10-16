package service

import (
	"context"

	"github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/models"
	pb "github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ch *ChatService) AddCommentService(ctx context.Context, req *pb.CommentRequest) (*pb.CommentResponse, error) {
	comment := models.Comment{
		ID:              primitive.NewObjectID().Hex(),
		ProblemID:       int(req.ProblemId),
		UserID:          req.UserId,
		Content:         req.Content,
		ParentCommentID: "",
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
		},
	}, nil
}

func (ch *ChatService) ReplyToCommentService(ctx context.Context, req *pb.ReplyRequest) (*pb.CommentResponse, error) {
	reply := models.Comment{
		ID:              primitive.NewObjectID().Hex(),
		UserID:          req.UserId,
		Content:         req.Content,
		ParentCommentID: req.CommentId,
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
		grpcComments = append(grpcComments, convertToGRPCComment(comment))
	}

	return &pb.FetchCommentsResponse{
		Comments: grpcComments,
	}, nil
}

func convertToGRPCComment(comment models.Comment) *pb.Comment {
	var grpcReplies []*pb.Comment
	for _, reply := range comment.Replies {
		grpcReplies = append(grpcReplies, convertToGRPCComment(reply))
	}

	return &pb.Comment{
		Id:              comment.ID,
		ProblemId:       uint32(comment.ProblemID),
		UserId:          comment.UserID,
		Content:         comment.Content,
		ParentCommentId: comment.ParentCommentID,
		Replies:         grpcReplies,
	}
}
