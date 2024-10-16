package interfaces

import pb "github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/proto"

type ChatServiceInter interface {
	// GetFriendChat(userid, friendid string, pagination models.Pagination) ([]models.Message, error)
	// MessageConsumer()

	CreateChatService(p *pb.Message) error
	FetchChatService(p *pb.ChatID) (*pb.ChatHistory, error)
}
