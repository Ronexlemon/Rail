package topics

import (
	"log"

	sharedKafka "github.com/ronexlemon/rail/shared/kafka"
	"github.com/segmentio/kafka-go"
)
 





var brokerURL = sharedKafka.GetEnv("KAFKA_BROKERS","kafka:9092")

var KafkaBrokerList = []string{brokerURL}

func initializeTopics() []string {
   
    authTopics := []string{TopicApiKeyRotation, TopicBusinessDeactivated, TopicUserCreated}
    businessTopics := []string{TopicClientCreated, TopicClientDeactivated}
    transactionTopics := []string{TopicCrossChainTransaction, TopicTransactionCreated, TopicTransactionUpdated} // Corrected name to be plural
    settlementTopics := []string{TopicSettlementProcessed, TopicSettlementFailed}
    walletTopics := []string{TopicBalanceUpdated, TopicWalletCreated}
    notificationTopics := []string{TopicNotificationCreated}

    
    allTopics := make([]string, 0)

    allTopics = append(allTopics, authTopics...)
    allTopics = append(allTopics, businessTopics...)
    allTopics = append(allTopics, transactionTopics...)
    allTopics = append(allTopics, settlementTopics...)
    allTopics = append(allTopics, walletTopics...)
    allTopics = append(allTopics, notificationTopics...)

    return allTopics
}

func RegisterTopics(){
	conn,err :=kafka.Dial("tcp",KafkaBrokerList[0])

	if err !=nil{
		log.Fatalf("Failed to Connect to Kafka %s",err)
	}
	defer conn.Close()

	topics := initializeTopics()

	for _,topic := range topics{
		err:=conn.CreateTopics(kafka.TopicConfig{
			Topic: topic,
			NumPartitions: 3,
			ReplicationFactor: 1,
		})
		if err != nil {
			log.Printf("Topic %s may already exist: %v", topic, err)
		} else {
			log.Printf("Topic %s registered successfully", topic)
		}
	}
	log.Println("✅ Kafka topics registration completed")
}
