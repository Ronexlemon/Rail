package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct{
	writer *kafka.Writer
}

func NewKafkaProducer(brokerURL,topic string)*KafkaProducer{
	writer:= kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerURL},
		Topic: topic,
		Balancer: &kafka.LeastBytes{}, // balances messages to partitions efficiently
	})
	return &KafkaProducer{
		writer: writer,
	}
}



func (p *KafkaProducer) Publish(key,value string)error{
	msg:= kafka.Message{
		Key: []byte(key),
		Value: []byte(value),
		Time: time.Now(),
	}
	err := p.writer.WriteMessages(context.Background(),msg)

	if err !=nil{
		log.Printf(" Kafka publish failed (topic =%s) : %v",p.writer.Topic,err)
		return err
	}
	log.Printf(" Published to kafka topic [%s]: %s",p.writer.Topic,value)
	return nil
}

func (p *KafkaProducer) Close()error{
	return p.writer.Close()
}