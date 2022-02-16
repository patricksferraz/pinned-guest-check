package event

type openGuestCheckMsg struct {
	GuestCheckID *string `json:"guest_check_id" valid:"uuid"`
	AttendantID  *string `json:"attendant_id" valid:"uuid"`
}

type OpenGuestCheck struct {
	Event `json:",inline"`
	Msg   *openGuestCheckMsg `json:"msg"`
}

type payGuestCheckMsg struct {
	GuestCheckID *string `json:"guest_check_id" valid:"uuid"`
}

type PayGuestCheck struct {
	Event `json:",inline"`
	Msg   *payGuestCheckMsg `json:"msg"`
}
