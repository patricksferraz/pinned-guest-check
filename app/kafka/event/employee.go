package event

type createEmployeeMsg struct {
	ID *string `json:"id" valid:"uuid"`
}

type Employee struct {
	Event `json:",inline" valid:"required"`
	Msg   *createEmployeeMsg `json:"msg" valid:"required"`
}
