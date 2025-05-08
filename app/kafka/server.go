package kafka

import (
	"fmt"

	"github.com/patricksferraz/pinned-guest-check/domain/service"
	"github.com/patricksferraz/pinned-guest-check/infra/client/kafka"
	"github.com/patricksferraz/pinned-guest-check/infra/db"
	"github.com/patricksferraz/pinned-guest-check/infra/repo"
)

func StartKafkaServer(orm *db.DbOrm, kc *kafka.KafkaConsumer, kp *kafka.KafkaProducer) {
	repository := repo.NewRepository(orm, kp)
	service := service.NewService(repository)

	fmt.Println("kafka pocessor has been started")
	processor := NewKafkaProcessor(service, kc)
	processor.Consume()
}
