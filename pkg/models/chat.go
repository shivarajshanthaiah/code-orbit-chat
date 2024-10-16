package models

import "time"

// type MessageReq struct {
// 	SenderID    string    `bson:"senderid"`
// 	RecipientID string    `bson:"recipientid"`
// 	Content     string    `bson:"content"`
// 	Timestamp   time.Time `bson:"timestamp"`
// }

// type Pagination struct {
// 	Limit  string
// 	OffSet string
// }

// type Message struct {
// 	ID          string    `bson:"_id"`
// 	SenderID    string    `bson:"senderid"`
// 	RecipientID string    `bson:"recipientid"`
// 	Content     string    `bson:"content"`
// 	Timestamp   time.Time `bson:"timestamp"`
// }

type Message struct {
	ID          string    `bson:"_id,omitempty"`
	SenderID    string    `json:"senderId"`
	RecipientID string    `json:"recipientId"`
	Content     string    `json:"content"`
	Timestamp   time.Time `json:"timestamp"`
}
