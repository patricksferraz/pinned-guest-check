package event

type createCustomerMsg struct {
	CustomerID *string `json:"customer_id" valid:"uuid"`
}

type Customer struct {
	Event `json:",inline"`
	Msg   *createCustomerMsg `json:"msg"`
}
