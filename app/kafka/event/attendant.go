package event

type createAttendantMsg struct {
	AttendantID *string `json:"attendant_id" valid:"uuid"`
}

type Attendant struct {
	Event `json:",inline"`
	Msg   *createAttendantMsg `json:"msg"`
}
