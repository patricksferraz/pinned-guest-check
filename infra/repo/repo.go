package repo

import (
	"context"
	"fmt"

	"github.com/c-4u/check-pad/domain/entity"
	"github.com/c-4u/check-pad/infra/db"
)

type Repository struct {
	Pg *db.PostgreSQL
}

func NewRepository(pg *db.PostgreSQL) *Repository {
	return &Repository{
		Pg: pg,
	}
}

func (r *Repository) CreateCustomer(ctx context.Context, customer *entity.Customer) error {
	err := r.Pg.Db.Create(customer).Error
	return err
}

func (r *Repository) FindCustomer(ctx context.Context, customerID *string) (*entity.Customer, error) {
	var e entity.Customer

	r.Pg.Db.First(&e, "id = ?", *customerID)

	if e.ID == nil {
		return nil, fmt.Errorf("no customer found")
	}

	return &e, nil
}

func (r *Repository) SaveCustomer(ctx context.Context, customer *entity.Customer) error {
	err := r.Pg.Db.Save(customer).Error
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

func (r *Repository) CreateCheckPad(ctx context.Context, checkPad *entity.CheckPad) error {
	err := r.Pg.Db.Create(checkPad).Error
	return err
}

func (r *Repository) FindCheckPad(ctx context.Context, checkPadID *string) (*entity.CheckPad, error) {
	var e entity.CheckPad
	r.Pg.Db.First(&e, "id = ?", *checkPadID)

	if e.ID == nil {
		return nil, fmt.Errorf("no check pad found")
	}

	return &e, nil
}

func (r *Repository) SaveCheckPad(ctx context.Context, checkPad *entity.CheckPad) error {
	err := r.Pg.Db.Save(checkPad).Error
	return err
}

func (r *Repository) FindCheckPadItem(ctx context.Context, checkPadItemID *string) (*entity.CheckPadItem, error) {
	var e entity.CheckPadItem
	r.Pg.Db.First(&e, "id = ?", *checkPadItemID)

	if e.ID == nil {
		return nil, fmt.Errorf("no check pad item found")
	}

	return &e, nil
}

func (r *Repository) SaveCheckPadItem(ctx context.Context, checkPadItem *entity.CheckPadItem) error {
	err := r.Pg.Db.Save(checkPadItem).Error
	return err
}
