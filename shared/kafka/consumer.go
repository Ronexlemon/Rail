package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct{
	reader *kafka.Reader
	topic string
	group  string
}

func NewKafkaConsumer(brokerURL,topic,groupID string)*KafkaConsumer{
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerURL},
		Topic: topic,
		GroupID: groupID,
		MaxBytes: 10e6, //10mb
		StartOffset: kafka.FirstOffset,
		MaxWait: 500 *time.Millisecond,
	})
	return &KafkaConsumer{
		reader: reader,
		topic: topic,
		group: groupID,
	}
}


// func (c *KafkaConsumer) Consume(ctx context.Context,handler func(key,value[]byte)){
// 	for{
// 		m, err := c.reader.FetchMessage(ctx)

// 		if err !=nil{
// 			if err == context.Canceled{
// 				log.Printf(" Kafka consumer for Topic [%s] stoped",c.topic)
// 			return

// 			}
// 			//log.Printf("Kafka consumer error %v",err)
// 			continue
			
// 		}
// 		log.Printf("Received message on [%s]: key= %s value = %s",c.topic,string(m.Key),string(m.Value))

// 		//Process the message
// 		handler(m.Key,m.Value)

// 		if err := c.reader.CommitMessages(ctx,m); err !=nil{
// 			log.Printf("Failed to commit message %v", err)
// 		}
// 	}
// }

func (c *KafkaConsumer) Consume(ctx context.Context, handler func(key, value []byte)) {
    for {
        m, err := c.reader.ReadMessage(ctx)
        if err != nil {
            if err == context.Canceled {
                log.Printf("Kafka consumer stopped for topic [%s]", c.topic)
                return
            }
            log.Printf("Kafka consumer error: %v", err)
            time.Sleep(time.Second)
            continue
        }

        log.Printf("Received message on [%s]: key=%s value=%s", c.topic, string(m.Key), string(m.Value))
        handler(m.Key, m.Value)
    }
}

func (c *KafkaConsumer) Close() error{
	log.Printf("Closing Kafka consumer for topic [%s]",c.topic)
	return c.reader.Close()
}