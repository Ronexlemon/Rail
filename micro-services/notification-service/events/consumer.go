package events

import (
	"context"
	"log"

	"github.com/ronexlemon/rail/micro-services/notification-service/configs"
	"github.com/ronexlemon/rail/shared/kafka"
	"github.com/ronexlemon/rail/shared/kafka/topics"
)

func ConsumeRegister() {
	brokerURL := configs.GetEnv("KAFKA_BROKERS", "kafka:9092")
	groupID := "business-service-group" 

	consumer := kafka.NewKafkaConsumer(brokerURL, topics.TopicUserCreated, groupID)

	ctx := context.Background()
	go consumer.Consume(ctx, func(key, value []byte) {
		log.Printf("Processing business.register event. Key=%s Value=%s\n", string(key), string(value))
		// TODO: unmarshal value into a struct and process business registration
	})

	
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Printf("Error closing consumer: %v", err)
		}
	}()
}
