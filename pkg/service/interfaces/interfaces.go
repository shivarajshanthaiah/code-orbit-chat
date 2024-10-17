package interfaces

import (
	"context"

	pb "github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/proto"
)

type ChatServiceInter interface {
	// GetFriendChat(userid, friendid string, pagination models.Pagination) ([]models.Message, error)
	// MessageConsumer()

	CreateChatService(p *pb.Message) error
	FetchChatService(p *pb.ChatID) (*pb.ChatHistory, error)

	AddCommentService(ctx context.Context, req *pb.CommentRequest) (*pb.CommentResponse, error)
	ReplyToCommentService(ctx context.Context, req *pb.ReplyRequest) (*pb.CommentResponse, error)
	GetCommentsForProblemService(ctx context.Context, req *pb.FetchCommentsRequest) (*pb.FetchCommentsResponse, error)

	GetUserCommentsService(ctx context.Context, req *pb.FetchUserCommentsRequest) (*pb.FetchUserCommentsResponse, error)
}
