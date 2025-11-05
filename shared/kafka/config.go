package kafka

import "os"

type KafkaConfig struct{
	BrokerUrl string
	Topic  string
}

func LoadKafkaConfig(brokerurl string ,topic string,topicFallback *string,brokerFallback *string)*KafkaConfig{
	return &KafkaConfig{
		BrokerUrl: getEnv(brokerurl,brokerFallback),
		Topic: getEnv(topic,topicFallback),
	}
}

func getEnv(key string,fallback *string)string{
	if val := os.Getenv(key); val !=""{
		return val
	}
	return *fallback
}