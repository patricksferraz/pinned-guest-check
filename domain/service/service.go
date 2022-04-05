package service

import (
	"context"

	"github.com/c-4u/pinned-guest-check/domain/entity"
	"github.com/c-4u/pinned-guest-check/domain/repo"
	"github.com/c-4u/pinned-guest-check/infra/client/kafka/topic"
	"github.com/c-4u/pinned-guest-check/utils"
)

type Service struct {
	Repo repo.RepoInterface
}

func NewService(repo repo.RepoInterface) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) CreateGuest(ctx context.Context, guestID *string) (*string, error) {
	guest, err := entity.NewGuest(guestID)
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreateGuest(ctx, guest); err != nil {
		return nil, err
	}

	return guest.ID, nil
}

func (s *Service) FindGuest(ctx context.Context, guestID *string) (*entity.Guest, error) {
	guest, err := s.Repo.FindGuest(ctx, guestID)
	if err != nil {
		return nil, err
	}

	return guest, nil
}

func (s *Service) CreatePlace(ctx context.Context, placeID *string) (*string, error) {
	place, err := entity.NewPlace(placeID)
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreatePlace(ctx, place); err != nil {
		return nil, err
	}

	return place.ID, nil
}

func (s *Service) FindPlace(ctx context.Context, placeID *string) (*entity.Place, error) {
	place, err := s.Repo.FindPlace(ctx, placeID)
	if err != nil {
		return nil, err
	}

	return place, nil
}

func (s *Service) CreateGuestCheck(ctx context.Context, local, guestID, placeID *string) (*string, error) {
	guest, err := s.Repo.FindGuest(ctx, guestID)
	if err != nil {
		return nil, err
	}

	place, err := s.Repo.FindPlace(ctx, placeID)
	if err != nil {
		return nil, err
	}

	guestCheck, err := entity.NewGuestCheck(local, guest, place)
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreateGuestCheck(ctx, guestCheck); err != nil {
		return nil, err
	}

	// TODO: adds retry
	event, err := entity.NewEvent(guestCheck)
	if err != nil {
		return nil, err
	}

	eMsg, err := event.ToJson()
	if err != nil {
		return nil, err
	}

	err = s.Repo.PublishEvent(ctx, utils.PString(topic.NEW_GUEST_CHECK), utils.PString(string(eMsg)), guestCheck.ID)
	if err != nil {
		return nil, err
	}

	return guestCheck.ID, nil
}

func (s *Service) FindGuestCheck(ctx context.Context, guestCheckID *string) (*entity.GuestCheck, error) {
	guestCheck, err := s.Repo.FindGuestCheck(ctx, guestCheckID)
	if err != nil {
		return nil, err
	}

	return guestCheck, nil
}

func (s *Service) WaitPaymentGuestCheck(ctx context.Context, guestCheckID *string) error {
	guestCheck, err := s.Repo.FindGuestCheck(ctx, guestCheckID)
	if err != nil {
		return err
	}

	err = guestCheck.WaitPayment()
	if err != nil {
		return err
	}

	if err = s.Repo.SaveGuestCheck(ctx, guestCheck); err != nil {
		return err
	}

	return nil
}

func (s *Service) CancelGuestCheck(ctx context.Context, guestCheckID, canceledReason *string) error {
	guestCheck, err := s.Repo.FindGuestCheck(ctx, guestCheckID)
	if err != nil {
		return err
	}

	err = guestCheck.Cancel(canceledReason)
	if err != nil {
		return err
	}

	if err = s.Repo.SaveGuestCheck(ctx, guestCheck); err != nil {
		return err
	}

	return nil
}

func (s *Service) OpenGuestCheck(ctx context.Context, guestCheckID, employeeID *string) error {
	guestCheck, err := s.Repo.FindGuestCheck(ctx, guestCheckID)
	if err != nil {
		return err
	}

	employee, err := s.Repo.FindEmployee(ctx, employeeID)
	if err != nil {
		return err
	}

	if err := guestCheck.Open(employee); err != nil {
		return err
	}

	if err = s.Repo.SaveGuestCheck(ctx, guestCheck); err != nil {
		return err
	}

	return nil
}

func (s *Service) PayGuestCheck(ctx context.Context, guestCheckID *string) error {
	guestCheck, err := s.Repo.FindGuestCheck(ctx, guestCheckID)
	if err != nil {
		return err
	}

	if err := guestCheck.Pay(); err != nil {
		return err
	}

	if err = s.Repo.SaveGuestCheck(ctx, guestCheck); err != nil {
		return err
	}

	return nil
}

func (s *Service) AddGuestCheckItem(ctx context.Context, guestCheckID, note *string, itemCode, quantity *int) (*string, error) {
	guestCheck, err := s.Repo.FindGuestCheck(ctx, guestCheckID)
	if err != nil {
		return nil, err
	}

	item, err := s.Repo.FindItemByCode(ctx, itemCode)
	if err != nil {
		return nil, err
	}

	guestCheckItem, err := entity.NewGuestCheckItem(item.Name, item.Code, quantity, item.Price, item.Discount, note, (*[]string)(item.Tags), guestCheck)
	if err != nil {
		return nil, err
	}

	// TODO: change strategy
	err = guestCheck.AddItem(guestCheckItem)
	if err != nil {
		return nil, err
	}

	// TODO: Adds transaction
	err = s.Repo.CreateGuestCheckItem(ctx, guestCheckItem)
	if err != nil {
		return nil, err
	}

	err = s.Repo.SaveGuestCheck(ctx, guestCheck)
	if err != nil {
		return nil, err
	}

	return guestCheckItem.ID, nil
}

func (s *Service) FindGuestCheckItem(ctx context.Context, guestCheckID, guestCheckItemID *string) (*entity.GuestCheckItem, error) {
	guestCheckItem, err := s.Repo.FindGuestCheckItem(ctx, guestCheckID, guestCheckItemID)
	if err != nil {
		return nil, err
	}

	return guestCheckItem, nil
}

func (s *Service) CancelGuestCheckItem(ctx context.Context, guestCheckID, guestCheckItemID, canceledReason *string) error {
	guestCheckItem, err := s.Repo.FindGuestCheckItem(ctx, guestCheckID, guestCheckItemID)
	if err != nil {
		return err
	}

	if err := guestCheckItem.Cancel(canceledReason); err != nil {
		return err
	}

	if err = s.Repo.SaveGuestCheckItem(ctx, guestCheckItem); err != nil {
		return err
	}

	return nil
}

func (s *Service) PrepareGuestCheckItem(ctx context.Context, guestCheckID, guestCheckItemID *string) error {
	guestCheckItem, err := s.Repo.FindGuestCheckItem(ctx, guestCheckID, guestCheckItemID)
	if err != nil {
		return err
	}

	if err := guestCheckItem.Prepare(); err != nil {
		return err
	}

	if err = s.Repo.SaveGuestCheckItem(ctx, guestCheckItem); err != nil {
		return err
	}

	return nil
}

func (s *Service) ReadyGuestCheckItem(ctx context.Context, guestCheckID, guestCheckItemID *string) error {
	guestCheckItem, err := s.Repo.FindGuestCheckItem(ctx, guestCheckID, guestCheckItemID)
	if err != nil {
		return err
	}

	if err := guestCheckItem.Ready(); err != nil {
		return err
	}

	if err = s.Repo.SaveGuestCheckItem(ctx, guestCheckItem); err != nil {
		return err
	}

	return nil
}

func (s *Service) ForwardGuestCheckItem(ctx context.Context, guestCheckID, guestCheckItemID *string) error {
	guestCheckItem, err := s.Repo.FindGuestCheckItem(ctx, guestCheckID, guestCheckItemID)
	if err != nil {
		return err
	}

	if err := guestCheckItem.Forward(); err != nil {
		return err
	}

	if err = s.Repo.SaveGuestCheckItem(ctx, guestCheckItem); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeliverGuestCheckItem(ctx context.Context, guestCheckID, guestCheckItemID *string) error {
	guestCheckItem, err := s.Repo.FindGuestCheckItem(ctx, guestCheckID, guestCheckItemID)
	if err != nil {
		return err
	}

	if err := guestCheckItem.Deliver(); err != nil {
		return err
	}

	if err = s.Repo.SaveGuestCheckItem(ctx, guestCheckItem); err != nil {
		return err
	}

	return nil
}

func (s *Service) CreateEmployee(ctx context.Context, employeeID *string) (*string, error) {
	employee, err := entity.NewEmployee(employeeID)
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreateEmployee(ctx, employee); err != nil {
		return nil, err
	}

	return employee.ID, nil
}

func (s *Service) FindEmployee(ctx context.Context, employeeID *string) (*entity.Employee, error) {
	employee, err := s.Repo.FindEmployee(ctx, employeeID)
	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (s *Service) CreateItem(ctx context.Context, id, name *string, code *int, price, discount *float64, available *bool, tags *[]string) (*string, error) {
	item, err := entity.NewItem(id, name, code, price, discount, available, tags)
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreateItem(ctx, item); err != nil {
		return nil, err
	}

	return item.ID, nil
}

func (s *Service) UpdateItem(ctx context.Context, itemID, name *string, available *bool, price, discount *float64, tags *[]string) error {
	item, err := s.Repo.FindItem(ctx, itemID)
	if err != nil {
		return err
	}

	if err = item.
		SetName(name).
		SetAvailable(available).
		SetPrice(price).
		SetDiscount(discount).
		SetTags(tags).IsValid(); err != nil {
		return err
	}

	if err = s.Repo.UpdateItem(ctx, item); err != nil {
		return err
	}

	return nil
}

func (s *Service) SearchGuestChecks(ctx context.Context, pageToken *string, pageSize *int) ([]*entity.GuestCheck, *string, error) {
	pagination, err := entity.NewPagination(pageToken, pageSize)
	if err != nil {
		return nil, nil, err
	}

	searchGuestChecks, err := entity.NewSearchGuestChecks(pagination)
	if err != nil {
		return nil, nil, err
	}

	guestChecks, nextPageToken, err := s.Repo.SearchGuestChecks(ctx, searchGuestChecks)
	if err != nil {
		return nil, nil, err
	}

	return guestChecks, nextPageToken, nil
}
