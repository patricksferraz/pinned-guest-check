package kafka

import (
	"context"
	"fmt"

	"github.com/c-4u/check-pad/app/kafka/event"
	"github.com/c-4u/check-pad/domain/service"
	"github.com/c-4u/check-pad/infra/client/kafka"
	"github.com/c-4u/check-pad/infra/client/kafka/topic"
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
	case topic.CANCEL_CHECK_PAD_ITEM:
		err := p.cancelCheckPadItem(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("cancel check pad item, error %s", err)
		}
	case topic.PREPARE_CHECK_PAD_ITEM:
		err := p.prepareCheckPadItem(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("prepare check pad item, error %s", err)
		}
	case topic.READY_CHECK_PAD_ITEM:
		err := p.readyCheckPadItem(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("ready check pad item, error %s", err)
		}
	case topic.FORWARD_CHECK_PAD_ITEM:
		err := p.forwardCheckPadItem(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("forward check pad item, error %s", err)
		}
	case topic.DELIVER_CHECK_PAD_ITEM:
		err := p.deliverCheckPadItem(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("deliver check pad item, error %s", err)
		}
	// CHECK_PAD
	case topic.OPEN_CHECK_PAD:
		err := p.openCheckPad(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("open check pad, error %s", err)
		}
	case topic.PAY_CHECK_PAD:
		err := p.payCheckPad(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("pay check pad, error %s", err)
		}
	// CUSTOMER
	case topic.NEW_CUSTOMER:
		err := p.createCustomer(msg)
		if err != nil {
			p.retry(msg)
			return fmt.Errorf("create customer, error %s", err)
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

func (p *KafkaProcessor) cancelCheckPadItem(msg *ckafka.Message) error {
	e := &event.CancelCheckPadItem{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	err = p.Service.CancelCheckPadItem(context.TODO(), e.Msg.CheckPadID, e.Msg.CheckPadItemID, e.Msg.CanceledReason)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) prepareCheckPadItem(msg *ckafka.Message) error {
	e := &event.PrepareCheckPadItem{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	err = p.Service.PrepareCheckPadItem(context.TODO(), e.Msg.CheckPadID, e.Msg.CheckPadItemID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) readyCheckPadItem(msg *ckafka.Message) error {
	e := &event.ReadyCheckPadItem{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	err = p.Service.ReadyCheckPadItem(context.TODO(), e.Msg.CheckPadID, e.Msg.CheckPadItemID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) forwardCheckPadItem(msg *ckafka.Message) error {
	e := &event.ForwardCheckPadItem{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	err = p.Service.ForwardCheckPadItem(context.TODO(), e.Msg.CheckPadID, e.Msg.CheckPadItemID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) deliverCheckPadItem(msg *ckafka.Message) error {
	e := &event.DeliverCheckPadItem{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	err = p.Service.DeliverCheckPadItem(context.TODO(), e.Msg.CheckPadID, e.Msg.CheckPadItemID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) openCheckPad(msg *ckafka.Message) error {
	e := &event.OpenCheckPad{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	err = p.Service.OpenCheckPad(context.TODO(), e.Msg.CheckPadID, e.Msg.AttendantID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) payCheckPad(msg *ckafka.Message) error {
	e := &event.PayCheckPad{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	err = p.Service.PayCheckPad(context.TODO(), e.Msg.CheckPadID)
	if err != nil {
		return err
	}

	return nil
}

func (p *KafkaProcessor) createCustomer(msg *ckafka.Message) error {
	e := &event.Customer{}
	err := e.ParseJson(msg.Value, e)
	if err != nil {
		return err
	}

	_, err = p.Service.CreateCustomer(context.TODO(), e.Msg.CustomerID)
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
