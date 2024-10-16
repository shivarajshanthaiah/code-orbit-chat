package service

// import (
// 	"context"
// 	"encoding/json"
// 	"log"
// 	"time"

// 	"github.com/IBM/sarama"
// 	"github.com/shivaraj-shanthaiah/code_orbit_chat/config"
// 	"github.com/shivaraj-shanthaiah/code_orbit_chat/pkg/models"
// 	"github.com/shivaraj-shanthaiah/code_orbit_chat/utils"
// )

// func (h ExampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
// 	return nil
// }

// func (h ExampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
// 	return nil
// }

// func (h ExampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
// 	for msg := range claim.Messages() {
// 		log.Printf("Received message: %s\n", string(msg.Value))
// 		chatMsg, err := h.Chatsvc.UnmarshalChatMessages(msg.Value)
// 		if err != nil {
// 			log.Printf("Error unmarshalling message: %v", err)
// 			continue
// 		}

// 		err = h.Chatsvc.Repo.StoreFriedsChat(*chatMsg)
// 		if err != nil {
// 			log.Printf("Error storing message in repository: %v", err)
// 			continue
// 		}
// 		sess.MarkMessage(msg, "")
// 	}
// 	return nil
// }

// func (ch *ChatService) MessageConsumer() {
// 	cfg := config.LoadConfig()
// 	if cfg == nil {
// 		log.Printf("Error loading config")
// 		return
// 	}

// 	brokers := []string{cfg.KafkaPort}
// 	topic := []string{"CHAT-TOPIC"}

// 	configs := sarama.NewConfig()
// 	configs.Version = sarama.V2_0_0_0
// 	configs.Consumer.Offsets.AutoCommit.Enable = true
// 	configs.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRange()
// 	configs.Consumer.Offsets.Initial = sarama.OffsetOldest

// 	log.Printf("Attempting to create consumer group with brokers: %v", brokers)

// 	admin, err := sarama.NewClusterAdmin(brokers, configs)
// 	if err != nil {
// 		log.Printf("Error creating cluster admin: %v", err)
// 	}
// 	defer admin.Close()

// 	err = admin.CreateTopic("CHAT-TOPIC", &sarama.TopicDetail{
// 		NumPartitions:     1,
// 		ReplicationFactor: 1,
// 	}, false)

// 	if err != nil {
// 		if err != sarama.ErrTopicAlreadyExists {
// 			log.Printf("Error creating topic: %v", err)
// 		}
// 	}

// 	consumerGroup, err := sarama.NewConsumerGroup(brokers, "chat-consumer-group", configs)
// 	if err != nil {
// 		log.Panicf("Error creating consumer group client: %v", err)
// 	}
// 	defer consumerGroup.Close()

// 	ctx := context.Background()
// 	handler := ExampleConsumerGroupHandler{Chatsvc: ch}

// 	log.Println("Starting to consume messages")

// 	for {
// 		err := consumerGroup.Consume(ctx, topic, handler)
// 		if err != nil {
// 			log.Printf("Error from consumer: %v", err)
// 			time.Sleep(time.Second * 5)
// 		}
// 	}

// }

// func (ch *ChatService) UnmarshalChatMessages(data []byte) (*models.MessageReq, error) {
// 	var message models.MessageReq
// 	err := json.Unmarshal(data, &message)
// 	if err != nil {
// 		return nil, err
// 	}
// 	message.Timestamp = time.Now()
// 	return &message, nil
// }

// func (ch *ChatService) GetFriendChat(userid, friendid string, pagination models.Pagination) ([]models.Message, error) {
// 	var err error
// 	pagination.OffSet, err = utils.Pagination(pagination.Limit, pagination.OffSet)
// 	if err != nil {
// 		return nil, err
// 	}
// 	_ = ch.Repo.UpdateReadAsMessages(userid, friendid)
// 	return ch.Repo.GetFriendChat(userid, friendid, pagination)
// }