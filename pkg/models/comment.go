package models

import "time"

type Comment struct {
	ID              string    `json:"id" bson:"_id,omitempty"`
	ProblemID       string    `json:"problem_id" bson:"problem_id"`
	UserID          string    `json:"user_id" bson:"user_id"`
	Content         string    `json:"content" bson:"content"`
	CreatedAt       time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" bson:"updated_at"`
	ParentCommentID string    `json:"parent_comment_id,omitempty" bson:"parent_comment_id,omitempty"`
	Replies         []Comment `json:"replies,omitempty" bson:"replies,omitempty"`
}
