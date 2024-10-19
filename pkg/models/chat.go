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
	ID          string    `bson:"_id,omitempty" json:"id,omitempty"`
	SenderID    string    `bson:"senderid" json:"senderid"`
	RecipientID string    `bson:"recipientid" json:"recipientid"`
	Content     string    `bson:"content" json:"content"`
	Timestamp   time.Time `bson:"timestamp" json:"timestamp"`
}
