package interfaces

import (
	"context"

	"github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/models"
)

type MongoRepoInter interface {
	// StoreFriedsChat(message models.MessageReq) error
	// UpdateReadAsMessages(userid, friendid string) error
	// GetFriendChat(userID, friendID string, pagination models.Pagination) ([]models.Message, error)

	Createchat(chat *models.Message) error
	Findchat(userID, receiverID string) (*[]models.Message, error)

	AddComment(ctx context.Context, comment *models.Comment) error
	AddReply(ctx context.Context, parentCommentID string, reply *models.Comment) error
	GetCommentsByProblemID(ctx context.Context, problemID int) ([]models.Comment, error)

	GetUserComments(ctx context.Context, userID string) ([]models.Comment, error)
}
