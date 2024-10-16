package repo

// import (
// 	"context"
// 	"strconv"
// 	"time"

// 	"github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/models"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func (r *MongoRepository) StoreFriedsChat(message models.MessageReq) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	_, err := r.Collection.InsertOne(ctx, message)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (r *MongoRepository) UpdateReadAsMessages(userid, friendid string) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	_, err := r.Collection.UpdateMany(ctx, bson.M{"senderid": bson.M{"$in": bson.A{friendid}}, "recipientid": bson.M{"$in": bson.A{userid}}}, bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: "send"}}}})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (r *MongoRepository) GetFriendChat(userID, friendID string, pagination models.Pagination) ([]models.Message, error) {

// 	var messages []models.Message
// 	filter := bson.M{"senderid": bson.M{"$in": bson.A{userID, friendID}}, "recipientid": bson.M{"$in": bson.A{friendID, userID}}}
// 	limit, _ := strconv.Atoi(pagination.Limit)
// 	offset, _ := strconv.Atoi(pagination.OffSet)
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	option := options.Find().SetLimit(int64(limit)).SetSkip(int64(offset))
// 	cursor, err := r.Collection.Find(ctx, filter, options.Find().SetSort(bson.D{{Key: "timestamp", Value: -1}}), option)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer cursor.Close(ctx)

// 	for cursor.Next(ctx) {
// 		var message models.Message
// 		if err := cursor.Decode(&message); err != nil {
// 			return nil, err
// 		}
// 		messages = append(messages, message)
// 	}
// 	return messages, nil
// }

// // func (r *MongoRepository) GetFriendChat(userID, friendID string) ([]models.Message, error) {
// // 	var messages []models.Message
// // 	filter := bson.M{"senderid": bson.M{"$in": bson.A{userID, friendID}}, "recipientid": bson.M{"$in": bson.A{friendID, userID}}}

// // 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// // 	defer cancel()

// // 	cursor, err := r.Collection.Find(ctx, filter, options.Find().SetSort(bson.D{{Key: "timestamp", Value: -1}}))
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	defer cursor.Close(ctx)
// // 	for cursor.Next(ctx) {
// // 		var message models.Message
// // 		if err := cursor.Decode(&message); err != nil {
// // 			return nil, err
// // 		}
// // 		messages = append(messages, message)
// // 	}
// // 	return messages, nil
// // }
