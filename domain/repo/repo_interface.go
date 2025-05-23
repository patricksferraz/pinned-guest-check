package repo

import (
	"context"

	"github.com/patricksferraz/pinned-guest-check/domain/entity"
)

type RepoInterface interface {
	CreateGuest(ctx context.Context, guest *entity.Guest) error
	FindGuest(ctx context.Context, guestID *string) (*entity.Guest, error)
	SaveGuest(ctx context.Context, guest *entity.Guest) error

	CreatePlace(ctx context.Context, place *entity.Place) error
	FindPlace(ctx context.Context, placeID *string) (*entity.Place, error)
	SavePlace(ctx context.Context, place *entity.Place) error

	CreateGuestCheck(ctx context.Context, guestCheck *entity.GuestCheck) error
	FindGuestCheck(ctx context.Context, guestCheckID *string) (*entity.GuestCheck, error)
	SaveGuestCheck(ctx context.Context, guestCheck *entity.GuestCheck) error
	SearchGuestChecks(ctx context.Context, searchGuestChecks *entity.SearchGuestChecks) ([]*entity.GuestCheck, *string, error)

	CreateGuestCheckItem(ctx context.Context, guestCheckItem *entity.GuestCheckItem) error
	FindGuestCheckItem(ctx context.Context, guestCheckID, guestCheckItemID *string) (*entity.GuestCheckItem, error)
	SaveGuestCheckItem(ctx context.Context, guestCheckItem *entity.GuestCheckItem) error

	CreateEmployee(ctx context.Context, employee *entity.Employee) error
	FindEmployee(ctx context.Context, employeeID *string) (*entity.Employee, error)
	SaveEmployee(ctx context.Context, employee *entity.Employee) error

	PublishEvent(ctx context.Context, topic, msg, key *string) error

	CreateItem(ctx context.Context, item *entity.Item) error
	FindItem(ctx context.Context, itemID *string) (*entity.Item, error)
	UpdateItem(ctx context.Context, item *entity.Item) error
	FindItemByCode(ctx context.Context, code *int) (*entity.Item, error)
}
