package event

type cancelGuestCheckItemMsg struct {
	GuestCheckID     *string `json:"guest_check_id" valid:"uuid"`
	GuestCheckItemID *string `json:"guest_check_item_id,omitempty" valid:"uuid"`
	CanceledReason   *string `json:"canceled_reason,omitempty"`
}

type CancelGuestCheckItem struct {
	Event `json:",inline"`
	Msg   *cancelGuestCheckItemMsg `json:"msg"`
}

type prepareGuestCheckItemMsg struct {
	GuestCheckID     *string `json:"guest_check_id" valid:"uuid"`
	GuestCheckItemID *string `json:"guest_check_item_id,omitempty" valid:"uuid"`
}

type PrepareGuestCheckItem struct {
	Event `json:",inline"`
	Msg   *prepareGuestCheckItemMsg `json:"msg"`
}

type readyGuestCheckItemMsg struct {
	GuestCheckID     *string `json:"guest_check_id" valid:"uuid"`
	GuestCheckItemID *string `json:"guest_check_item_id,omitempty" valid:"uuid"`
}

type ReadyGuestCheckItem struct {
	Event `json:",inline"`
	Msg   *readyGuestCheckItemMsg `json:"msg"`
}

type forwardGuestCheckItemMsg struct {
	GuestCheckID     *string `json:"guest_check_id" valid:"uuid"`
	GuestCheckItemID *string `json:"guest_check_item_id,omitempty" valid:"uuid"`
}

type ForwardGuestCheckItem struct {
	Event `json:",inline"`
	Msg   *forwardGuestCheckItemMsg `json:"msg"`
}

type deliverGuestCheckItemMsg struct {
	GuestCheckID     *string `json:"guest_check_id" valid:"uuid"`
	GuestCheckItemID *string `json:"guest_check_item_id,omitempty" valid:"uuid"`
}

type DeliverGuestCheckItem struct {
	Event `json:",inline"`
	Msg   *deliverGuestCheckItemMsg `json:"msg"`
}
