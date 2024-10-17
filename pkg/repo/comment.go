package repo

import (
	"context"
	"log"

	"github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *MongoRepository) AddComment(ctx context.Context, comment *models.Comment) error {
	_, err := r.CommentCollection.InsertOne(ctx, comment)
	return err
}

func (r *MongoRepository) AddReply(ctx context.Context, parentCommentID string, reply *models.Comment) error {
	filter := bson.M{"_id": parentCommentID}
	update := bson.M{
		"$push": bson.M{"replies": reply},
	}

	_, err := r.CommentCollection.UpdateOne(ctx, filter, update)
	return err
}

func (r *MongoRepository) GetCommentsByProblemID(ctx context.Context, problemID int) ([]models.Comment, error) {
	var comments []models.Comment

	filter := bson.M{"problem_id": problemID, "parent_comment_id": bson.M{"$exists": false}} // Root comments only
	cursor, err := r.CommentCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var comment models.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (r *MongoRepository) GetUserComments(ctx context.Context, userID string) ([]models.Comment, error) {
	var comments []models.Comment

	// Find all comments where either:
	// 1. The main comment is by the user
	// 2. Any reply in the replies array is by the user
	filter := bson.M{
		"$or": []bson.M{
			{"user_id": userID},
			{"replies.user_id": userID},
		},
	}

	log.Printf("Fetching comments with filter: %+v", filter)

	cursor, err := r.CommentCollection.Find(ctx, filter)
	if err != nil {
		log.Printf("Error finding comments: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var comment models.Comment
		if err := cursor.Decode(&comment); err != nil {
			log.Printf("Error decoding comment: %v", err)
			return nil, err
		}

		log.Printf("Found comment: ID=%s, Content=%s", comment.ID, comment.Content)
		if len(comment.Replies) > 0 {
			log.Printf("Comment has %d replies", len(comment.Replies))
		}

		comments = append(comments, comment)
	}

	return comments, nil
}
