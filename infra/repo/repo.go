package repo

import (
	"context"
	"fmt"

	"github.com/c-4u/pinned-guest-check/domain/entity"
	"github.com/c-4u/pinned-guest-check/infra/client/kafka"
	"github.com/c-4u/pinned-guest-check/infra/db"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Repository struct {
	Pg *db.PostgreSQL
	Kp *kafka.KafkaProducer
}

func NewRepository(pg *db.PostgreSQL, kp *kafka.KafkaProducer) *Repository {
	return &Repository{
		Pg: pg,
		Kp: kp,
	}
}

func (r *Repository) CreateGuest(ctx context.Context, guest *entity.Guest) error {
	err := r.Pg.Db.Create(guest).Error
	return err
}

func (r *Repository) FindGuest(ctx context.Context, guestID *string) (*entity.Guest, error) {
	var e entity.Guest

	r.Pg.Db.First(&e, "id = ?", *guestID)

	if e.ID == nil {
		return nil, fmt.Errorf("no guest found")
	}

	return &e, nil
}

func (r *Repository) SaveGuest(ctx context.Context, guest *entity.Guest) error {
	err := r.Pg.Db.Save(guest).Error
	return err
}

func (r *Repository) CreatePlace(ctx context.Context, place *entity.Place) error {
	err := r.Pg.Db.Create(place).Error
	return err
}

func (r *Repository) FindPlace(ctx context.Context, placeID *string) (*entity.Place, error) {
	var e entity.Place
	r.Pg.Db.First(&e, "id = ?", *placeID)

	if e.ID == nil {
		return nil, fmt.Errorf("no place found")
	}

	return &e, nil
}

func (r *Repository) SavePlace(ctx context.Context, place *entity.Place) error {
	err := r.Pg.Db.Save(place).Error
	return err
}

func (r *Repository) CreateGuestCheck(ctx context.Context, guestCheck *entity.GuestCheck) error {
	err := r.Pg.Db.Create(guestCheck).Error
	return err
}

func (r *Repository) FindGuestCheck(ctx context.Context, guestCheckID *string) (*entity.GuestCheck, error) {
	var e entity.GuestCheck
	r.Pg.Db.Preload("Items").First(&e, "id = ?", *guestCheckID)

	if e.ID == nil {
		return nil, fmt.Errorf("no guest check found")
	}

	return &e, nil
}

func (r *Repository) SaveGuestCheck(ctx context.Context, guestCheck *entity.GuestCheck) error {
	// TODO: infinity loop when saving guest check with backoff Items
	err := r.Pg.Db.Save(guestCheck).Error
	return err
}

func (r *Repository) CreateGuestCheckItem(ctx context.Context, guestCheckItem *entity.GuestCheckItem) error {
	err := r.Pg.Db.Create(guestCheckItem).Error
	return err
}

func (r *Repository) FindGuestCheckItem(ctx context.Context, guestCheckID, guestCheckItemID *string) (*entity.GuestCheckItem, error) {
	var e entity.GuestCheckItem
	r.Pg.Db.Preload("GuestCheck").First(&e, "id = ? AND guest_check_id = ?", *guestCheckItemID, *guestCheckID)

	if e.ID == nil {
		return nil, fmt.Errorf("no guest check item found")
	}

	return &e, nil
}

func (r *Repository) SaveGuestCheckItem(ctx context.Context, guestCheckItem *entity.GuestCheckItem) error {
	err := r.Pg.Db.Save(guestCheckItem).Error
	return err
}

func (r *Repository) CreateEmployee(ctx context.Context, employee *entity.Employee) error {
	err := r.Pg.Db.Create(employee).Error
	return err
}

func (r *Repository) FindEmployee(ctx context.Context, employeeID *string) (*entity.Employee, error) {
	var e entity.Employee

	r.Pg.Db.First(&e, "id = ?", *employeeID)

	if e.ID == nil {
		return nil, fmt.Errorf("no employee found")
	}

	return &e, nil
}

func (r *Repository) SaveEmployee(ctx context.Context, employee *entity.Employee) error {
	err := r.Pg.Db.Save(employee).Error
	return err
}

func (r *Repository) PublishEvent(ctx context.Context, topic, msg, key *string) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: topic, Partition: ckafka.PartitionAny},
		Value:          []byte(*msg),
		Key:            []byte(*key),
	}
	err := r.Kp.Producer.Produce(message, r.Kp.DeliveryChan)
	if err != nil {
		return err
	}
	return nil
}
