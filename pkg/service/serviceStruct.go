package service

import (
	interRepo "github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/repo/interfaces"
	"github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/service/interfaces"
)

type ChatService struct {
	Repo interRepo.MongoRepoInter
}

func NewChatService(repo interRepo.MongoRepoInter) interfaces.ChatServiceInter {
	return &ChatService{
		Repo: repo,
	}
}

// type ExampleConsumerGroupHandler struct {
// 	Chatsvc *ChatService
// }
