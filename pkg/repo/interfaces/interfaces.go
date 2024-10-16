package interfaces

import "github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/models"

type MongoRepoInter interface {
	// StoreFriedsChat(message models.MessageReq) error
	// UpdateReadAsMessages(userid, friendid string) error
	// GetFriendChat(userID, friendID string, pagination models.Pagination) ([]models.Message, error)

	Createchat(chat *models.Message) error
	Findchat(userID, receiverID string) (*[]models.Message, error)
}
