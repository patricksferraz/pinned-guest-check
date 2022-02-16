package event

type createGuestMsg struct {
	GuestID *string `json:"guest_id" valid:"uuid"`
}

type Guest struct {
	Event `json:",inline"`
	Msg   *createGuestMsg `json:"msg"`
}
