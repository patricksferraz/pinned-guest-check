package event

type cancelGuestCheckItemMsg struct {
	GuestCheckID     *string `json:"guest_check_id" valid:"uuid"`
	GuestCheckItemID *string `json:"guest_check_item_id,omitempty" valid:"uuid"`
	CanceledReason   *string `json:"canceled_reason,omitempty"`
}

type CancelGuestCheckItem struct {
	Event `json:",inline" valid:"required"`
	Msg   *cancelGuestCheckItemMsg `json:"msg" valid:"required"`
}

type prepareGuestCheckItemMsg struct {
	GuestCheckID     *string `json:"guest_check_id" valid:"uuid"`
	GuestCheckItemID *string `json:"guest_check_item_id,omitempty" valid:"uuid"`
}

type PrepareGuestCheckItem struct {
	Event `json:",inline" valid:"required"`
	Msg   *prepareGuestCheckItemMsg `json:"msg" valid:"required"`
}

type readyGuestCheckItemMsg struct {
	GuestCheckID     *string `json:"guest_check_id" valid:"uuid"`
	GuestCheckItemID *string `json:"guest_check_item_id,omitempty" valid:"uuid"`
}

type ReadyGuestCheckItem struct {
	Event `json:",inline" valid:"required"`
	Msg   *readyGuestCheckItemMsg `json:"msg" valid:"required"`
}

type forwardGuestCheckItemMsg struct {
	GuestCheckID     *string `json:"guest_check_id" valid:"uuid"`
	GuestCheckItemID *string `json:"guest_check_item_id,omitempty" valid:"uuid"`
}

type ForwardGuestCheckItem struct {
	Event `json:",inline" valid:"required"`
	Msg   *forwardGuestCheckItemMsg `json:"msg" valid:"required"`
}

type deliverGuestCheckItemMsg struct {
	GuestCheckID     *string `json:"guest_check_id" valid:"uuid"`
	GuestCheckItemID *string `json:"guest_check_item_id,omitempty" valid:"uuid"`
}

type DeliverGuestCheckItem struct {
	Event `json:",inline" valid:"required"`
	Msg   *deliverGuestCheckItemMsg `json:"msg" valid:"required"`
}
