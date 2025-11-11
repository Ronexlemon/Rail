package topics

import (
	"log"

	sharedKafka "github.com/ronexlemon/rail/shared/kafka"
	"github.com/segmentio/kafka-go"
)
 





var brokerURL = sharedKafka.GetEnv("KAFKA_BROKERS","kafka:9092")

var KafkaBrokerList = []string{"localhost:9092"}

func initializeTopics() []string {

    authTopics := []string{TopicBusinessRegister, TopicBusinessVerified,TopicBusinessApiKeyGenerated,TopicBusinessapiKeyRevoked,TopicBusinessDeactivated}
    businessTopics := []string{TopicBusinessRegistered, TopicBusinessPendingVerification,TopicBusinessWalletLinked,TopicBusinessSuspended,TopicCustomerCreated,TopicCustomerDeactivated,TopicCustomerDeactivated}
    transactionTopics := []string{TopicTransactionInitiated,TopicTransactionApproved, TopicTransactionCancelled, TopicTransactionFailed,TopicTransactionValidated } // Corrected name to be plural
    settlementTopics := []string{TopicSettlementCompleted, TopicSettlementFailed,TopicSettlementInProgress,TopicSettlementInitiated,TopicSettlementInProgress,TopicOnchainTransactionConfirmed,TopicOnchainTransactionFailed}
    walletTopics := []string{TopicWalletCreated,TopicWalletFunded,TopicWalletDebited,TopicwalletCredited,TopicWalletFunded,TopicWalletSuspended, TopicWalletBalanceUpdated,TopicWalletClosed}
    notificationTopics := []string{TopicNotificationCreated,TopicNotificationSent}
	compliance :=[]string{TopicKYCRejected,TopicBusinessFlagged}
	report :=[]string{TopicReportGenerated,TopicAnalyticUpdated}


    
    allTopics := make([]string, 0)

    allTopics = append(allTopics, authTopics...)
    allTopics = append(allTopics, businessTopics...)
    allTopics = append(allTopics, transactionTopics...)
    allTopics = append(allTopics, settlementTopics...)
    allTopics = append(allTopics, walletTopics...)
    allTopics = append(allTopics, notificationTopics...)
	allTopics = append(allTopics, compliance...)
	allTopics = append(allTopics, report...)

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
