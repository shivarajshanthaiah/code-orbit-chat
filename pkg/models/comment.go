package models

type Comment struct {
	ID              string    `json:"id" bson:"_id,omitempty"`
	ProblemID       int       `json:"problem_id" bson:"problem_id"`
	UserID          string    `json:"user_id" bson:"user_id"`
	Content         string    `json:"content" bson:"content"`
	ParentCommentID string    `json:"parent_comment_id,omitempty" bson:"parent_comment_id,omitempty"`
	Replies         []Comment `json:"replies,omitempty" bson:"replies,omitempty"`
}
