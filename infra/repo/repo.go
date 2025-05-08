package repo

import (
	"context"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/patricksferraz/pinned-guest-check/domain/entity"
	"github.com/patricksferraz/pinned-guest-check/infra/client/kafka"
	"github.com/patricksferraz/pinned-guest-check/infra/db"
)

type Repository struct {
	Orm *db.DbOrm
	Kp  *kafka.KafkaProducer
}

func NewRepository(orm *db.DbOrm, kp *kafka.KafkaProducer) *Repository {
	return &Repository{
		Orm: orm,
		Kp:  kp,
	}
}

func (r *Repository) CreateGuest(ctx context.Context, guest *entity.Guest) error {
	err := r.Orm.Db.Create(guest).Error
	return err
}

func (r *Repository) FindGuest(ctx context.Context, guestID *string) (*entity.Guest, error) {
	var e entity.Guest

	r.Orm.Db.First(&e, "id = ?", *guestID)

	if e.ID == nil {
		return nil, fmt.Errorf("no guest found")
	}

	return &e, nil
}

func (r *Repository) SaveGuest(ctx context.Context, guest *entity.Guest) error {
	err := r.Orm.Db.Save(guest).Error
	return err
}

func (r *Repository) CreatePlace(ctx context.Context, place *entity.Place) error {
	err := r.Orm.Db.Create(place).Error
	return err
}

func (r *Repository) FindPlace(ctx context.Context, placeID *string) (*entity.Place, error) {
	var e entity.Place
	r.Orm.Db.First(&e, "id = ?", *placeID)

	if e.ID == nil {
		return nil, fmt.Errorf("no place found")
	}

	return &e, nil
}

func (r *Repository) SavePlace(ctx context.Context, place *entity.Place) error {
	err := r.Orm.Db.Save(place).Error
	return err
}

func (r *Repository) CreateGuestCheck(ctx context.Context, guestCheck *entity.GuestCheck) error {
	err := r.Orm.Db.Create(guestCheck).Error
	return err
}

func (r *Repository) FindGuestCheck(ctx context.Context, guestCheckID *string) (*entity.GuestCheck, error) {
	var e entity.GuestCheck
	r.Orm.Db.Preload("Items").First(&e, "id = ?", *guestCheckID)

	if e.ID == nil {
		return nil, fmt.Errorf("no guest check found")
	}

	return &e, nil
}

func (r *Repository) SaveGuestCheck(ctx context.Context, guestCheck *entity.GuestCheck) error {
	// TODO: infinity loop when saving guest check with backoff Items
	err := r.Orm.Db.Save(guestCheck).Error
	return err
}

func (r *Repository) CreateGuestCheckItem(ctx context.Context, guestCheckItem *entity.GuestCheckItem) error {
	err := r.Orm.Db.Create(guestCheckItem).Error
	return err
}

func (r *Repository) FindGuestCheckItem(ctx context.Context, guestCheckID, guestCheckItemID *string) (*entity.GuestCheckItem, error) {
	var e entity.GuestCheckItem
	r.Orm.Db.Preload("GuestCheck").First(&e, "id = ? AND guest_check_id = ?", *guestCheckItemID, *guestCheckID)

	if e.ID == nil {
		return nil, fmt.Errorf("no guest check item found")
	}

	return &e, nil
}

func (r *Repository) SaveGuestCheckItem(ctx context.Context, guestCheckItem *entity.GuestCheckItem) error {
	err := r.Orm.Db.Save(guestCheckItem).Error
	return err
}

func (r *Repository) CreateEmployee(ctx context.Context, employee *entity.Employee) error {
	err := r.Orm.Db.Create(employee).Error
	return err
}

func (r *Repository) FindEmployee(ctx context.Context, employeeID *string) (*entity.Employee, error) {
	var e entity.Employee

	r.Orm.Db.First(&e, "id = ?", *employeeID)

	if e.ID == nil {
		return nil, fmt.Errorf("no employee found")
	}

	return &e, nil
}

func (r *Repository) SaveEmployee(ctx context.Context, employee *entity.Employee) error {
	err := r.Orm.Db.Save(employee).Error
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

func (r *Repository) CreateItem(ctx context.Context, item *entity.Item) error {
	err := r.Orm.Db.Create(item).Error
	return err
}

func (r *Repository) FindItem(ctx context.Context, itemID *string) (*entity.Item, error) {
	var e entity.Item

	r.Orm.Db.First(&e, "id = ?", *itemID)

	if e.ID == nil {
		return nil, fmt.Errorf("no item found")
	}

	return &e, nil
}

func (r *Repository) UpdateItem(ctx context.Context, item *entity.Item) error {
	err := r.Orm.Db.Updates(item).Error
	return err
}

func (r *Repository) FindItemByCode(ctx context.Context, code *int) (*entity.Item, error) {
	var e entity.Item

	r.Orm.Db.First(&e, "code = ?", *code)

	if e.ID == nil {
		return nil, fmt.Errorf("no item found")
	}

	return &e, nil
}

func (r *Repository) SearchGuestChecks(ctx context.Context, searchGuestChecks *entity.SearchGuestChecks) ([]*entity.GuestCheck, *string, error) {
	var e []*entity.GuestCheck

	q := r.Orm.Db
	if searchGuestChecks.PageToken != nil {
		q = q.Where("token < ?", *searchGuestChecks.PageToken)
	}
	err := q.Order("token DESC").
		Limit(*searchGuestChecks.PageSize).
		Find(&e).Error
	if err != nil {
		return nil, nil, err
	}

	if len(e) < *searchGuestChecks.PageSize {
		return e, nil, nil
	}

	return e, e[len(e)-1].Token, nil
}
