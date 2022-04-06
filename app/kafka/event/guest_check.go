package event

type openGuestCheckMsg struct {
	ID         *string `json:"id" valid:"uuid"`
	AttendedBy *string `json:"attended_by" valid:"uuid"`
}

type OpenGuestCheck struct {
	Event `json:",inline" valid:"required"`
	Msg   *openGuestCheckMsg `json:"msg" valid:"required"`
}
