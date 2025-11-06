package config

import (
    "encoding/json"
    "log"

    "github.com/ronexlemon/rail/shared/kafka"
)

func PublishEvent(key string, value interface{}) {
    brokerUrl := GetEnv("KAFKA_BROKERS", "kafka:9092")

    producer := kafka.NewKafkaProducer(brokerUrl, "business.register")

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
