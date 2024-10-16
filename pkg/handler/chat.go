package handler

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/models"
	pb "github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/proto"
)

type MessageHandle struct {
	MQue []models.Message
	mu   sync.Mutex
}

var messageHandleObject = MessageHandle{}

func (ch *ChatServiceServer) Connect(csi pb.ChatService_ConnectServer) error {

	errch := make(chan error)

	go ch.receiveFromStream(csi, errch)

	go sendToStream(csi, errch)

	return <-errch
}

func (ch *ChatServiceServer) receiveFromStream(csi pb.ChatService_ConnectServer, errch chan error) {

	for {
		msg, err := csi.Recv()
		if err != nil {
			log.Printf("error in receiving msg from client : %v", err)
			errch <- err
		} else {
			messageHandleObject.mu.Lock()
			ch.svc.CreateChatService(msg)
			messageHandleObject.MQue = append(messageHandleObject.MQue, models.Message{
				SenderID:    msg.User_ID,
				RecipientID: msg.Receiver_ID,
				Content:     msg.Content,
			})
			messageHandleObject.mu.Unlock()

			// log.Printf("%v",messageHandleObject.MQue[len(messageHandleObject.MQue)-1])

		}
	}
}

func sendToStream(csi pb.ChatService_ConnectServer, errch chan error) {

	for {

		for {
			time.Sleep(500 * time.Millisecond)

			messageHandleObject.mu.Lock()

			if len(messageHandleObject.MQue) == 0 {
				messageHandleObject.mu.Unlock()
				break
			}

			userID := messageHandleObject.MQue[0].SenderID
			recieverID := messageHandleObject.MQue[0].RecipientID
			message := messageHandleObject.MQue[0].Content

			messageHandleObject.mu.Unlock()

			if userID != recieverID {
				// fmt.Println(message)
				err := csi.Send(&pb.Message{
					User_ID:     userID,
					Receiver_ID: recieverID,
					Content:     message,
				})

				if err != nil {
					errch <- err
				}

				messageHandleObject.mu.Lock()

				if len(messageHandleObject.MQue) > 1 {
					messageHandleObject.MQue = messageHandleObject.MQue[1:]
				} else {
					messageHandleObject.MQue = []models.Message{}
				}

				messageHandleObject.mu.Unlock()
			}
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func (ch *ChatServiceServer) FetchHistory(ctx context.Context, p *pb.ChatID) (*pb.ChatHistory, error) {
	response, err := ch.svc.FetchChatService(p)
	if err != nil {
		return response, err
	}
	return response, nil

}
