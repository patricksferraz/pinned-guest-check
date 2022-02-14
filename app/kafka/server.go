package kafka

import (
	"fmt"

	"github.com/c-4u/check-pad/domain/service"
	"github.com/c-4u/check-pad/infra/client/kafka"
	"github.com/c-4u/check-pad/infra/db"
	"github.com/c-4u/check-pad/infra/repo"
)

func StartKafkaServer(pg *db.PostgreSQL, kc *kafka.KafkaConsumer, kp *kafka.KafkaProducer) {
	repository := repo.NewRepository(pg, kp)
	service := service.NewService(repository)

	fmt.Println("kafka pocessor has been started")
	processor := NewKafkaProcessor(service, kc)
	processor.Consume()
}
