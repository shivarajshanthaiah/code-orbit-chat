package di

import (
	"log"

	"github.com/shivaraj-shanthaiah/code_orbit_chat/config"
	"github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/db"
	"github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/handler"
	"github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/repo"
	"github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/server"
	"github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/service"
)

func Init() {
	cnfg := config.LoadConfig()

	mongoClient, err := db.ConnectMongoDB(cnfg)
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}

	mongoDB := mongoClient.Database(cnfg.DBname)

	chatRepo := repo.NewMongoRepository(mongoDB)
	chatService := service.NewChatService(chatRepo)
	chatHandler := handler.NewChatServiceServer(chatService)

	err = server.NewGrpcChatServer(cnfg.GrpcPort, chatHandler)
	if err != nil {
		log.Fatalf("failed to start gRPC server %v", err.Error())
	}
}
