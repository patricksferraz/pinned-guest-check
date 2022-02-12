package rest

import "time"

type Base struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type HTTPResponse struct {
	Msg string `json:"msg,omitempty" example:"any message"`
}

type IDResponse struct {
	ID string `json:"id"`
}

type Customer struct {
	Base `json:",inline"`
}

type Place struct {
	Base `json:",inline"`
}

type CreateCheckPadRequest struct {
	Local      string `json:"local"`
	CustomerID string `json:"customer_id"`
	PlaceID    string `json:"place_id"`
}

type CheckPad struct {
	Base           `json:",inline"`
	TotalPrice     float64 `json:"total_price,omitempty"`
	TotalDiscount  float64 `json:"total_discount,omitempty"`
	FinalPrice     float64 `json:"final_price,omitempty"`
	Status         int     `json:"status"`
	CanceledReason string  `json:"canceled_reason,omitempty"`
	Local          string  `json:"local"`
	CustomerID     string  `json:"customer_id"`
	PlaceID        string  `json:"place_id"`
}

type CancelCheckPadRequest struct {
	CanceledReason string `json:"canceled_reason"`
}

type AddCheckPadItemRequest struct {
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	Discount  float64 `json:"discount,omitempty"`
	Note      string  `json:"note,omitempty"`
	Tag       string  `json:"tag"`
}

type CancelCheckPadItemRequest struct {
	CanceledReason string `json:"canceled_reason"`
}

type CheckPadItem struct {
	Base           `json:",inline"`
	Status         int     `json:"status"`
	CanceledReason string  `json:"canceled_reason,omitempty"`
	Name           string  `json:"name"`
	Quantity       int     `json:"quantity"`
	UnitPrice      float64 `json:"unit_price"`
	Discount       float64 `json:"discount,omitempty"`
	TotalPrice     float64 `json:"total_price,omitempty"`
	FinalPrice     float64 `json:"final_price"`
	Note           string  `json:"note,omitempty"`
	Tag            string  `json:"tag"`
	CheckPadID     string  `json:"check_pad_id"`
}
