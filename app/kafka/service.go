package kafka

import (
	"context"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/patricksferraz/pinned-guest-check/app/kafka/event"
	"github.com/patricksferraz/pinned-guest-check/domain/service"
	"github.com/patricksferraz/pinned-guest-check/infra/client/kafka"
	"github.com/patricksferraz/pinned-guest-check/infra/client/kafka/topic"
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
	// MENU ITEM
	case topic.NEW_MENU_ITEM:
		err := p.createItem(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("create item, error %s", err)
		}
	case topic.UPDATE_MENU_ITEM:
		err := p.updateItem(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("update item, error %s", err)
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

	err = p.Service.OpenGuestCheck(context.TODO(), e.Msg.ID, e.Msg.AttendedBy)
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

func (p *KafkaProcessor) createItem(msg *ckafka.Message) error {
	e := &event.Item{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	_, err = p.Service.CreateItem(context.TODO(), e.Msg.ID, e.Msg.Name, e.Msg.Code, e.Msg.Price, e.Msg.Discount, e.Msg.Available, e.Msg.Tags)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) updateItem(msg *ckafka.Message) error {
	e := &event.Item{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	err = p.Service.UpdateItem(context.TODO(), e.Msg.ID, e.Msg.Name, e.Msg.Available, e.Msg.Price, e.Msg.Discount, e.Msg.Tags)
	if err != nil {
		return err
	}

	return nil
}
