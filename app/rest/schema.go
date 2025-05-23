package rest

import "time"

type Base struct {
	ID        *string    `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type HTTPResponse struct {
	Msg *string `json:"msg,omitempty" example:"any message"`
}

type IDResponse struct {
	ID *string `json:"id"`
}

type Guest struct {
	Base `json:",inline"`
}

type Place struct {
	Base `json:",inline"`
}

type CreateGuestCheckRequest struct {
	Local   *string `json:"local"`
	GuestID *string `json:"guest_id"`
	PlaceID *string `json:"place_id"`
}

type GuestCheck struct {
	Base           `json:",inline"`
	TotalPrice     *float64 `json:"total_price,omitempty"`
	TotalDiscount  *float64 `json:"total_discount,omitempty"`
	FinalPrice     *float64 `json:"final_price,omitempty"`
	Status         *int     `json:"status"`
	CanceledReason *string  `json:"canceled_reason,omitempty"`
	Local          *string  `json:"local"`
	GuestID        *string  `json:"guest_id"`
	PlaceID        *string  `json:"place_id"`
}

type CancelGuestCheckRequest struct {
	CanceledReason *string `json:"canceled_reason"`
}

type AddGuestCheckItemRequest struct {
	ItemCode *int    `json:"item_code"`
	Quantity *int    `json:"quantity"`
	Note     *string `json:"note,omitempty"`
}

type CancelGuestCheckItemRequest struct {
	CanceledReason *string `json:"canceled_reason"`
}

type GuestCheckItem struct {
	Base           `json:",inline"`
	Status         *int     `json:"status"`
	CanceledReason *string  `json:"canceled_reason,omitempty"`
	Name           *string  `json:"name"`
	Code           *int     `json:"code"`
	Quantity       *int     `json:"quantity"`
	UnitPrice      *float64 `json:"unit_price"`
	Discount       *float64 `json:"discount,omitempty"`
	TotalPrice     *float64 `json:"total_price,omitempty"`
	FinalPrice     *float64 `json:"final_price"`
	Note           *string  `json:"note,omitempty"`
	Tag            *string  `json:"tag"`
	GuestCheckID   *string  `json:"guest_check_id"`
}

type Employee struct {
	Base `json:",inline"`
}

type SearchGuestChecksRequest struct {
	PageToken *string `json:"page_token" query:"page_token"`
	PageSize  *int    `json:"page_size" query:"page_size"`
}

type SearchGuestChecksResponse struct {
	GuestChecks   []*GuestCheck `json:"guest_checks"`
	NextPageToken *string       `json:"next_page_token"`
}

type PayGuestCheckRequest struct {
	Tip *float64 `json:"tip"`
}
