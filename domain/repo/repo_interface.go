package repo

import (
	"context"

	"github.com/c-4u/check-pad/domain/entity"
)

type RepoInterface interface {
	CreateCustomer(ctx context.Context, customer *entity.Customer) error
	FindCustomer(ctx context.Context, customerID *string) (*entity.Customer, error)
	SaveCustomer(ctx context.Context, customer *entity.Customer) error

	CreatePlace(ctx context.Context, place *entity.Place) error
	FindPlace(ctx context.Context, placeID *string) (*entity.Place, error)
	SavePlace(ctx context.Context, place *entity.Place) error

	CreateCheckPad(ctx context.Context, checkPad *entity.CheckPad) error
	FindCheckPad(ctx context.Context, checkPadID *string) (*entity.CheckPad, error)
	SaveCheckPad(ctx context.Context, checkPad *entity.CheckPad) error

	CreateCheckPadItem(ctx context.Context, checkPadItem *entity.CheckPadItem) error
	FindCheckPadItem(ctx context.Context, checkPadID, checkPadItemID *string) (*entity.CheckPadItem, error)
	SaveCheckPadItem(ctx context.Context, checkPadItem *entity.CheckPadItem) error
}
