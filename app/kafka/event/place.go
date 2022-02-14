package event

type createPlaceMsg struct {
	PlaceID *string `json:"place_id" valid:"uuid"`
}

type Place struct {
	Event `json:",inline"`
	Msg   *createPlaceMsg `json:"msg"`
}
