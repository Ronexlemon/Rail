package events

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	"github.com/ronexlemon/rail/micro-services/business-service/configs"
	"github.com/ronexlemon/rail/micro-services/business-service/internal/repository"
	"github.com/ronexlemon/rail/micro-services/business-service/internal/service"
	"github.com/ronexlemon/rail/shared/kafka"
	"github.com/ronexlemon/rail/shared/kafka/topics"
)
type UserModel struct {
	ID        string    `json:"id"`
	Name      string `json:"name"`
	Email     string    `json:"email"`
	Role     string `json:"role"`
	Password  string    `json:"-"` // hidden in JSON
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}


func ConsumeRegister() {
	brokerURL := configs.GetEnv("KAFKA_BROKERS", "kafka:9092")
	groupID := "business-service-group"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-sig:
			log.Println("Shutting down business-service...")
			cancel()
			return
		default:
			err := startConsumer(ctx, brokerURL, groupID)
			if err != nil {
				log.Printf("Consumer crashed: %v. Retrying in 5s...\n", err)
				time.Sleep(5 * time.Second)
			}
		}
	}
}

func startConsumer(ctx context.Context, brokerURL, groupID string) error {
	c := kafka.NewKafkaConsumer(brokerURL, topics.TopicBusinessRegistered, groupID)
	defer c.Close()

	log.Printf("Kafka consumer connected (topic=%s, group=%s)", topics.TopicBusinessRegistered, groupID)

	errChan := make(chan error, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("🔥 Panic in consumer: %v\n%s", r, string(debug.Stack()))
				errChan <- r.(error)
			}
		}()

		c.Consume(ctx, func(key, value []byte) {
			log.Printf(" Processing event [user-created] Key=%s Value=%s", string(key), string(value))

			var payload UserModel
			if err := json.Unmarshal(value, &payload); err != nil {
				log.Printf("Failed to unmarshal payload: %v", err)
				return
			}

			// TODO: process payload here (DB updates, etc.) just a place holder for producer
            
			//PublishEvent(string(key),value)
			repo:= repository.NewBusinessRepository()
			service:= service.NewBusinesssService(repo)
			service.CreateBusinessService(payload.Name,payload.Email,payload.ID)
		})
	}()

	select {
	case <-ctx.Done():
		return nil
	case err := <-errChan:
		return err
	}
}
