package kafka

import (
	"context"
	"fmt"

	"github.com/c-4u/guest-check/app/kafka/event"
	"github.com/c-4u/guest-check/domain/service"
	"github.com/c-4u/guest-check/infra/client/kafka"
	"github.com/c-4u/guest-check/infra/client/kafka/topic"
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
	// CHEKC_PAD_ITEM
	case topic.CANCEL_GUEST_CHECK_ITEM:
		err := p.cancelGuestCheckItem(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("cancel guest check item, error %s", err)
		}
	case topic.PREPARE_GUEST_CHECK_ITEM:
		err := p.prepareGuestCheckItem(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("prepare guest check item, error %s", err)
		}
	case topic.READY_GUEST_CHECK_ITEM:
		err := p.readyGuestCheckItem(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("ready guest check item, error %s", err)
		}
	case topic.FORWARD_GUEST_CHECK_ITEM:
		err := p.forwardGuestCheckItem(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("forward guest check item, error %s", err)
		}
	case topic.DELIVER_GUEST_CHECK_ITEM:
		err := p.deliverGuestCheckItem(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("deliver guest check item, error %s", err)
		}
	// GUEST_CHECK
	case topic.OPEN_GUEST_CHECK:
		err := p.openGuestCheck(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("open guest check, error %s", err)
		}
	case topic.PAY_GUEST_CHECK:
		err := p.payGuestCheck(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("pay guest check, error %s", err)
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
	// ATTENDANT
	case topic.NEW_ATTENDANT:
		err := p.createAttendant(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("create attendant, error %s", err)
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

func (p *KafkaProcessor) cancelGuestCheckItem(msg *ckafka.Message) error {
	e := &event.CancelGuestCheckItem{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	err = p.Service.CancelGuestCheckItem(context.TODO(), e.Msg.GuestCheckID, e.Msg.GuestCheckItemID, e.Msg.CanceledReason)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) prepareGuestCheckItem(msg *ckafka.Message) error {
	e := &event.PrepareGuestCheckItem{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	err = p.Service.PrepareGuestCheckItem(context.TODO(), e.Msg.GuestCheckID, e.Msg.GuestCheckItemID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) readyGuestCheckItem(msg *ckafka.Message) error {
	e := &event.ReadyGuestCheckItem{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	err = p.Service.ReadyGuestCheckItem(context.TODO(), e.Msg.GuestCheckID, e.Msg.GuestCheckItemID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) forwardGuestCheckItem(msg *ckafka.Message) error {
	e := &event.ForwardGuestCheckItem{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	err = p.Service.ForwardGuestCheckItem(context.TODO(), e.Msg.GuestCheckID, e.Msg.GuestCheckItemID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) deliverGuestCheckItem(msg *ckafka.Message) error {
	e := &event.DeliverGuestCheckItem{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	err = p.Service.DeliverGuestCheckItem(context.TODO(), e.Msg.GuestCheckID, e.Msg.GuestCheckItemID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) openGuestCheck(msg *ckafka.Message) error {
	e := &event.OpenGuestCheck{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	err = p.Service.OpenGuestCheck(context.TODO(), e.Msg.GuestCheckID, e.Msg.AttendantID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) payGuestCheck(msg *ckafka.Message) error {
	e := &event.PayGuestCheck{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	err = p.Service.PayGuestCheck(context.TODO(), e.Msg.GuestCheckID)
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

	_, err = p.Service.CreateGuest(context.TODO(), e.Msg.GuestID)
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

	_, err = p.Service.CreatePlace(context.TODO(), e.Msg.PlaceID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) createAttendant(msg *ckafka.Message) error {
	e := &event.Attendant{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	_, err = p.Service.CreateAttendant(context.TODO(), e.Msg.AttendantID)
	if err != nil {
		return err
	}

	return nil
}
