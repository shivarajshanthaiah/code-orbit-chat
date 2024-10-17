package utils

import (
	"github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/models"
	pb "github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/proto"
)

func ConvertToGRPCCommentForProblem(comment models.Comment) *pb.Comment {
	var grpcReplies []*pb.Comment
	for _, reply := range comment.Replies {
		grpcReplies = append(grpcReplies, ConvertToGRPCCommentForProblem(reply))
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

func ConvertToGRPCCommentForComment(comment models.Comment) *pb.Comment {
	grpcComment := &pb.Comment{
		Id:              comment.ID,
		ProblemId:       uint32(comment.ProblemID),
		UserId:          comment.UserID,
		Content:         comment.Content,
		ParentCommentId: comment.ParentCommentID,
	}

	// Convert replies if they exist
	if len(comment.Replies) > 0 {
		grpcComment.Replies = make([]*pb.Comment, 0, len(comment.Replies))
		for _, reply := range comment.Replies {
			grpcComment.Replies = append(grpcComment.Replies, ConvertToGRPCCommentForComment(reply))
		}
	}

	return grpcComment
}
