package service

import (
	"context"

	"github.com/c-4u/check-pad/domain/entity"
	"github.com/c-4u/check-pad/domain/repo"
)

type Service struct {
	Repo repo.RepoInterface
}

func NewService(repo repo.RepoInterface) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) CreateCustomer(ctx context.Context) (*string, error) {
	customer, err := entity.NewCustomer()
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreateCustomer(ctx, customer); err != nil {
		return nil, err
	}

	return customer.ID, nil
}

func (s *Service) FindCustomer(ctx context.Context, customerID *string) (*entity.Customer, error) {
	customer, err := s.Repo.FindCustomer(ctx, customerID)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *Service) CreatePlace(ctx context.Context) (*string, error) {
	place, err := entity.NewPlace()
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

func (s *Service) CreateCheckPad(ctx context.Context, local, customerID, placeID *string) (*string, error) {
	customer, err := s.Repo.FindCustomer(ctx, customerID)
	if err != nil {
		return nil, err
	}

	place, err := s.Repo.FindPlace(ctx, placeID)
	if err != nil {
		return nil, err
	}

	checkPad, err := entity.NewCheckPad(local, customer, place)
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreateCheckPad(ctx, checkPad); err != nil {
		return nil, err
	}

	return checkPad.ID, nil
}

func (s *Service) FindCheckPad(ctx context.Context, checkPadID *string) (*entity.CheckPad, error) {
	checkPad, err := s.Repo.FindCheckPad(ctx, checkPadID)
	if err != nil {
		return nil, err
	}

	return checkPad, nil
}

func (s *Service) ReopenCheckPad(ctx context.Context, checkPadID *string) error {
	checkPad, err := s.Repo.FindCheckPad(ctx, checkPadID)
	if err != nil {
		return err
	}

	if err := checkPad.Reopen(); err != nil {
		return err
	}

	if err = s.Repo.SaveCheckPad(ctx, checkPad); err != nil {
		return err
	}

	return nil
}

func (s *Service) WaitPaymentCheckPad(ctx context.Context, checkPadID *string) error {
	checkPad, err := s.Repo.FindCheckPad(ctx, checkPadID)
	if err != nil {
		return err
	}

	if err := checkPad.WaitPayment(); err != nil {
		return err
	}

	if err = s.Repo.SaveCheckPad(ctx, checkPad); err != nil {
		return err
	}

	return nil
}

func (s *Service) CancelCheckPad(ctx context.Context, checkPadID, canceledReason *string) error {
	checkPad, err := s.Repo.FindCheckPad(ctx, checkPadID)
	if err != nil {
		return err
	}

	if err := checkPad.Cancel(canceledReason); err != nil {
		return err
	}

	if err = s.Repo.SaveCheckPad(ctx, checkPad); err != nil {
		return err
	}

	return nil
}

func (s *Service) CancelCheckPadItem(ctx context.Context, checkPadItemID, canceledReason *string) error {
	checkPadItem, err := s.Repo.FindCheckPadItem(ctx, checkPadItemID)
	if err != nil {
		return err
	}

	if err := checkPadItem.Cancel(canceledReason); err != nil {
		return err
	}

	if err = s.Repo.SaveCheckPadItem(ctx, checkPadItem); err != nil {
		return err
	}

	return nil
}

func (s *Service) PrepareCheckPadItem(ctx context.Context, checkPadItemID, canceledReason *string) error {
	checkPadItem, err := s.Repo.FindCheckPadItem(ctx, checkPadItemID)
	if err != nil {
		return err
	}

	if err := checkPadItem.Prepare(); err != nil {
		return err
	}

	if err = s.Repo.SaveCheckPadItem(ctx, checkPadItem); err != nil {
		return err
	}

	return nil
}

func (s *Service) ForwardCheckPadItem(ctx context.Context, checkPadItemID, canceledReason *string) error {
	checkPadItem, err := s.Repo.FindCheckPadItem(ctx, checkPadItemID)
	if err != nil {
		return err
	}

	if err := checkPadItem.Forward(); err != nil {
		return err
	}

	if err = s.Repo.SaveCheckPadItem(ctx, checkPadItem); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeliverCheckPadItem(ctx context.Context, checkPadItemID, canceledReason *string) error {
	checkPadItem, err := s.Repo.FindCheckPadItem(ctx, checkPadItemID)
	if err != nil {
		return err
	}

	if err := checkPadItem.Deliver(); err != nil {
		return err
	}

	if err = s.Repo.SaveCheckPadItem(ctx, checkPadItem); err != nil {
		return err
	}

	return nil
}
