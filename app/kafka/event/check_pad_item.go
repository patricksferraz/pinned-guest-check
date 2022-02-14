package event

type cancelCheckPadItemMsg struct {
	CheckPadID     *string `json:"check_pad_id" valid:"uuid"`
	CheckPadItemID *string `json:"check_pad_item_id,omitempty" valid:"uuid"`
	CanceledReason *string `json:"canceled_reason,omitempty"`
}

type CancelCheckPadItem struct {
	Event `json:",inline"`
	Msg   *cancelCheckPadItemMsg `json:"msg"`
}

type prepareCheckPadItemMsg struct {
	CheckPadID     *string `json:"check_pad_id" valid:"uuid"`
	CheckPadItemID *string `json:"check_pad_item_id,omitempty" valid:"uuid"`
}

type PrepareCheckPadItem struct {
	Event `json:",inline"`
	Msg   *prepareCheckPadItemMsg `json:"msg"`
}

type forwardCheckPadItemMsg struct {
	CheckPadID     *string `json:"check_pad_id" valid:"uuid"`
	CheckPadItemID *string `json:"check_pad_item_id,omitempty" valid:"uuid"`
}

type ForwardCheckPadItem struct {
	Event `json:",inline"`
	Msg   *forwardCheckPadItemMsg `json:"msg"`
}

type deliverCheckPadItemMsg struct {
	CheckPadID     *string `json:"check_pad_id" valid:"uuid"`
	CheckPadItemID *string `json:"check_pad_item_id,omitempty" valid:"uuid"`
}

type DeliverCheckPadItem struct {
	Event `json:",inline"`
	Msg   *deliverCheckPadItemMsg `json:"msg"`
}
