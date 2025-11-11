package kafka

import "os"

type KafkaConfig struct{
	BrokerUrl string
	Topic  string
}

func LoadKafkaConfig(brokerurl string ,topic string,topicFallback string,brokerFallback string)*KafkaConfig{
	return &KafkaConfig{
		BrokerUrl: GetEnv(brokerurl,brokerFallback),
		Topic: GetEnv(topic,topicFallback),
	}
}

func GetEnv(key string,fallback string)string{
	if val := os.Getenv(key); val !=""{
		return val
	}
	return fallback
}

