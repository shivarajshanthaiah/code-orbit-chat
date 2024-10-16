package repo

import (
	"context"
	"time"

	"github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *MongoRepository) Createchat(chat *models.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.Collection.InsertOne(ctx, chat)
	if err != nil {
		return err
	}
	return nil
}

// Findchat method finds the chat from MongoDB using user and receiver id
func (r *MongoRepository) Findchat(userID, receiverID string) (*[]models.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var chat []models.Message
	filter := bson.M{"senderId": userID, "recipientId": receiverID}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &chat); err != nil {
		return nil, err
	}

	return &chat, nil
}
