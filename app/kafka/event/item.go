package event

type createItemMsg struct {
	ID        *string   `json:"id" valid:"uuid"`
	Code      *int      `json:"code" valid:"-"`
	Name      *string   `json:"name" valid:"required"`
	Available *bool     `json:"available" valid:"-"`
	Price     *float64  `json:"price" valid:"required"`
	Discount  *float64  `json:"discount" valid:"-"`
	Tags      *[]string `json:"tags" valid:"-"`
}

type Item struct {
	Event `json:",inline" valid:"required"`
	Msg   *createItemMsg `json:"msg" valid:"required"`
}
