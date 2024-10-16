package repo

import (
	inter "github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/repo/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	Collection        *mongo.Collection
	CommentCollection *mongo.Collection
}

func NewMongoRepository(mongo *mongo.Database) inter.MongoRepoInter {
	return &MongoRepository{
		Collection:        mongo.Collection("messages"),
		CommentCollection: mongo.Collection("comments"),
	}
}

