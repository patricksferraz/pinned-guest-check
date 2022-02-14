package service

import (
	"context"

	"github.com/c-4u/check-pad/domain/entity"
	"github.com/c-4u/check-pad/domain/repo"
	"github.com/c-4u/check-pad/infra/client/kafka/topic"
	"github.com/c-4u/check-pad/utils"
)

type Service struct {
	Repo repo.RepoInterface
}

func NewService(repo repo.RepoInterface) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) CreateCustomer(ctx context.Context, customerID *string) (*string, error) {
	customer, err := entity.NewCustomer(customerID)
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

	// TODO: adds retry
	event, err := entity.NewEvent(checkPad)
	if err != nil {
		return nil, err
	}

	eMsg, err := event.ToJson()
	if err != nil {
		return nil, err
	}

	err = s.Repo.PublishEvent(ctx, utils.PString(topic.NEW_CHECK_PAD), utils.PString(string(eMsg)), checkPad.ID)
	if err != nil {
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

func (s *Service) WaitPaymentCheckPad(ctx context.Context, checkPadID *string) error {
	checkPad, err := s.Repo.FindCheckPad(ctx, checkPadID)
	if err != nil {
		return err
	}

	wpMsg, err := checkPad.WaitPayment()
	if err != nil {
		return err
	}

	if err = s.Repo.SaveCheckPad(ctx, checkPad); err != nil {
		return err
	}

	// TODO: adds retry
	event, err := entity.NewEvent(wpMsg)
	if err != nil {
		return err
	}

	eMsg, err := event.ToJson()
	if err != nil {
		return err
	}

	err = s.Repo.PublishEvent(ctx, utils.PString(topic.WAIT_PAYMENT_CHECK_PAD), utils.PString(string(eMsg)), checkPad.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) CancelCheckPad(ctx context.Context, checkPadID, canceledReason *string) error {
	checkPad, err := s.Repo.FindCheckPad(ctx, checkPadID)
	if err != nil {
		return err
	}

	cMsg, err := checkPad.Cancel(canceledReason)
	if err != nil {
		return err
	}

	if err = s.Repo.SaveCheckPad(ctx, checkPad); err != nil {
		return err
	}

	// TODO: adds retry
	event, err := entity.NewEvent(cMsg)
	if err != nil {
		return err
	}

	eMsg, err := event.ToJson()
	if err != nil {
		return err
	}

	err = s.Repo.PublishEvent(ctx, utils.PString(topic.CANCEL_CHECK_PAD), utils.PString(string(eMsg)), checkPad.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) OpenCheckPad(ctx context.Context, checkPadID, attendantID *string) error {
	checkPad, err := s.Repo.FindCheckPad(ctx, checkPadID)
	if err != nil {
		return err
	}

	attendant, err := s.Repo.FindAttendant(ctx, attendantID)
	if err != nil {
		return err
	}

	if err := checkPad.SetAttendant(attendant); err != nil {
		return err
	}

	if err = s.Repo.SaveCheckPad(ctx, checkPad); err != nil {
		return err
	}

	return nil
}

func (s *Service) PayCheckPad(ctx context.Context, checkPadID *string) error {
	checkPad, err := s.Repo.FindCheckPad(ctx, checkPadID)
	if err != nil {
		return err
	}

	if err := checkPad.Pay(); err != nil {
		return err
	}

	if err = s.Repo.SaveCheckPad(ctx, checkPad); err != nil {
		return err
	}

	return nil
}

func (s *Service) AddCheckPadItem(ctx context.Context, name *string, code, quantity *int, unitPrice *float64, discount *float64, note *string, tag *string, checkPadID *string) (*string, error) {
	checkPad, err := s.Repo.FindCheckPad(ctx, checkPadID)
	if err != nil {
		return nil, err
	}

	checkPadItem, err := entity.NewCheckPadItem(name, code, quantity, unitPrice, discount, note, tag, checkPad)
	if err != nil {
		return nil, err
	}

	err = checkPad.AddItem(checkPadItem)
	if err != nil {
		return nil, err
	}

	// TODO: Adds transaction
	err = s.Repo.CreateCheckPadItem(ctx, checkPadItem)
	if err != nil {
		return nil, err
	}

	err = s.Repo.SaveCheckPad(ctx, checkPad)
	if err != nil {
		return nil, err
	}

	// TODO: adds retry
	event, err := entity.NewEvent(checkPadItem)
	if err != nil {
		return nil, err
	}

	eMsg, err := event.ToJson()
	if err != nil {
		return nil, err
	}

	err = s.Repo.PublishEvent(ctx, utils.PString(topic.NEW_CHECK_PAD_ITEM), utils.PString(string(eMsg)), checkPad.ID)
	if err != nil {
		return nil, err
	}

	return checkPadItem.ID, nil
}

func (s *Service) FindCheckPadItem(ctx context.Context, checkPadID, checkPadItemID *string) (*entity.CheckPadItem, error) {
	checkPadItem, err := s.Repo.FindCheckPadItem(ctx, checkPadID, checkPadItemID)
	if err != nil {
		return nil, err
	}

	return checkPadItem, nil
}

func (s *Service) CancelCheckPadItem(ctx context.Context, checkPadID, checkPadItemID, canceledReason *string) error {
	checkPadItem, err := s.Repo.FindCheckPadItem(ctx, checkPadID, checkPadItemID)
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

func (s *Service) PrepareCheckPadItem(ctx context.Context, checkPadID, checkPadItemID *string) error {
	checkPadItem, err := s.Repo.FindCheckPadItem(ctx, checkPadID, checkPadItemID)
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

func (s *Service) ForwardCheckPadItem(ctx context.Context, checkPadID, checkPadItemID *string) error {
	checkPadItem, err := s.Repo.FindCheckPadItem(ctx, checkPadID, checkPadItemID)
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

func (s *Service) DeliverCheckPadItem(ctx context.Context, checkPadID, checkPadItemID *string) error {
	checkPadItem, err := s.Repo.FindCheckPadItem(ctx, checkPadID, checkPadItemID)
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

func (s *Service) CreateAttendant(ctx context.Context, attendantID *string) (*string, error) {
	attendant, err := entity.NewAttendant(attendantID)
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreateAttendant(ctx, attendant); err != nil {
		return nil, err
	}

	return attendant.ID, nil
}

func (s *Service) FindAttendant(ctx context.Context, attendantID *string) (*entity.Attendant, error) {
	attendant, err := s.Repo.FindAttendant(ctx, attendantID)
	if err != nil {
		return nil, err
	}

	return attendant, nil
}
