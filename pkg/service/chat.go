package service

import (
	"time"

	"github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/models"
	pb "github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/proto"
)

func (ch *ChatService) CreateChatService(p *pb.Message) error {
	chat := &models.Message{
		SenderID:    p.User_ID,
		RecipientID: p.Receiver_ID,
		Content:     p.Content,
		Timestamp:   time.Now(),
	}
	err := ch.Repo.Createchat(chat)
	if err != nil {
		return err
	}

	return nil
}

// func (ch *ChatService) FetchChatService(p *pb.ChatID) (*pb.ChatHistory, error) {

// 	userHistory, err := ch.Repo.Findchat(p.User_ID, p.Receiver_ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	receiverHistory, err := ch.Repo.Findchat(p.Receiver_ID, p.User_ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var chats []*pb.Message
// 	for _, msg := range *userHistory {
// 		chats = append(chats, &pb.Message{
// 			Chat_ID:     msg.ID,
// 			User_ID:     msg.SenderID,
// 			Receiver_ID: msg.RecipientID,
// 			Content:     msg.Content,
// 		})
// 	}
// 	for _, msg := range *receiverHistory {
// 		chats = append(chats, &pb.Message{
// 			Chat_ID:     msg.ID,
// 			User_ID:     msg.SenderID,
// 			Receiver_ID: msg.RecipientID,
// 			Content:     msg.Content,
// 		})
// 	}
// 	sortByChatID(chats)
// 	return &pb.ChatHistory{
// 		Chats: chats,
// 	}, nil
// }

// func sortByChatID(chats []*pb.Message) {
// 	sort.Slice(chats, func(i, j int) bool {
// 		return chats[i].Chat_ID < chats[j].Chat_ID
// 	})
// }

func (ch *ChatService) FetchChatService(p *pb.ChatID) (*pb.ChatHistory, error) {

	userHistory, err := ch.Repo.Findchat(p.User_ID, p.Receiver_ID)
	if err != nil {
		return nil, err
	}
	receiverHistory, err := ch.Repo.Findchat(p.Receiver_ID, p.User_ID)
	if err != nil {
		return nil, err
	}
	var chats []*pb.Message
	for _, msg := range *userHistory {
		chats = append(chats, &pb.Message{
			Chat_ID:     msg.ID,
			User_ID:     msg.SenderID,
			Receiver_ID: msg.RecipientID,
			Content:     msg.Content,
			Timestamp:   msg.Timestamp.String(), // Assuming Timestamp is a field in pb.Message
		})
	}
	for _, msg := range *receiverHistory {
		chats = append(chats, &pb.Message{
			Chat_ID:     msg.ID,
			User_ID:     msg.SenderID,
			Receiver_ID: msg.RecipientID,
			Content:     msg.Content,
			Timestamp:   msg.Timestamp.String(),
		})
	}

	return &pb.ChatHistory{
		Chats: chats,
	}, nil
}
