package kafka

import (
	"context"
	"fmt"

	"github.com/c-4u/pinned-guest-check/app/kafka/event"
	"github.com/c-4u/pinned-guest-check/domain/service"
	"github.com/c-4u/pinned-guest-check/infra/client/kafka"
	"github.com/c-4u/pinned-guest-check/infra/client/kafka/topic"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProcessor struct {
	Service *service.Service
	Kc      *kafka.KafkaConsumer
}

func NewKafkaProcessor(service *service.Service, kafkaConsumer *kafka.KafkaConsumer) *KafkaProcessor {
	return &KafkaProcessor{
		Service: service,
		Kc:      kafkaConsumer,
	}
}

func (p *KafkaProcessor) Consume() {
	p.Kc.Consumer.SubscribeTopics(p.Kc.ConsumerTopics, nil)
	for {
		msg, err := p.Kc.Consumer.ReadMessage(-1)
		if err == nil {
			// fmt.Println(string(msg.Value))
			err := p.processMessage(msg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (p *KafkaProcessor) processMessage(msg *ckafka.Message) error {
	switch _topic := *msg.TopicPartition.Topic; _topic {
	// GUEST_CHECK
	case topic.OPEN_GUEST_CHECK:
		err := p.openGuestCheck(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("open guest check, error %s", err)
		}
	// GUEST
	case topic.NEW_GUEST:
		err := p.createGuest(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("create guest, error %s", err)
		}
	// PLACE
	case topic.NEW_PLACE:
		err := p.createPlace(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("create place, error %s", err)
		}
	// EMPLOYEE
	case topic.NEW_EMPLOYEE:
		err := p.createEmployee(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("create employee, error %s", err)
		}
	default:
		return fmt.Errorf("not a valid topic %s", string(msg.Value))
	}

	return nil
}

func (p *KafkaProcessor) retry(msg *ckafka.Message) error {
	err := p.Kc.Consumer.Seek(ckafka.TopicPartition{
		Topic:     msg.TopicPartition.Topic,
		Partition: msg.TopicPartition.Partition,
		Offset:    msg.TopicPartition.Offset,
	}, -1)

	return err
}

func (p *KafkaProcessor) openGuestCheck(msg *ckafka.Message) error {
	e := &event.OpenGuestCheck{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	err = p.Service.OpenGuestCheck(context.TODO(), e.Msg.GuestCheckID, e.Msg.EmployeeID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) createGuest(msg *ckafka.Message) error {
	e := &event.Guest{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	_, err = p.Service.CreateGuest(context.TODO(), e.Msg.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) createPlace(msg *ckafka.Message) error {
	e := &event.Place{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	_, err = p.Service.CreatePlace(context.TODO(), e.Msg.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) createEmployee(msg *ckafka.Message) error {
	e := &event.Employee{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	_, err = p.Service.CreateEmployee(context.TODO(), e.Msg.ID)
	if err != nil {
		return err
	}

	return nil
}
