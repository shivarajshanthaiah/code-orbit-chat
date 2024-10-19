package repo

import (
	"context"
	"time"

	"github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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


func (r *MongoRepository) Findchat(userID, receiverID string) (*[]models.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"$or": []bson.M{
			{"senderid": userID, "recipientid": receiverID},
			{"senderid": receiverID, "recipientid": userID},
		},
	}
	opts := options.Find().SetSort(bson.D{{Key: "timestamp", Value: 1}})
	cursor, err := r.Collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var messages []models.Message
	if err := cursor.All(ctx, &messages); err != nil {
		return nil, err
	}

	return &messages, nil
}