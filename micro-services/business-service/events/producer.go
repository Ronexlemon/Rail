package events

import (
	"encoding/json"
	"log"

	"github.com/ronexlemon/rail/micro-services/business-service/configs"
	"github.com/ronexlemon/rail/shared/kafka"
	"github.com/ronexlemon/rail/shared/kafka/topics"
)

func PublishEvent(key string, value interface{}) {
    brokerUrl := configs.GetEnv("KAFKA_BROKERS", "kafka:9092")
    
    
    producer := kafka.NewKafkaProducer(brokerUrl, topics.TopicWalletCreated)

    // Convert value to JSON
    jsonValue, err := json.Marshal(value)
    if err != nil {
        log.Printf("Failed to serialize event value: %v\n", err)
        return
    }

    // Publish
    if err := producer.Publish(key, string(jsonValue)); err != nil {
        log.Printf("Failed to publish event: %v\n", err)
    }
}
