package event

type openCheckPadMsg struct {
	CheckPadID  *string `json:"check_pad_id" valid:"uuid"`
	AttendantID *string `json:"attendant_id" valid:"uuid"`
}

type OpenCheckPad struct {
	Event `json:",inline"`
	Msg   *openCheckPadMsg `json:"msg"`
}

type payCheckPadMsg struct {
	CheckPadID *string `json:"check_pad_id" valid:"uuid"`
}

type PayCheckPad struct {
	Event `json:",inline"`
	Msg   *payCheckPadMsg `json:"msg"`
}
