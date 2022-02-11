package rest

import (
	"time"
)

type Base struct {
	ID        string    `json:"id,omitempty" binding:"uuid"`
	CreatedAt time.Time `json:"created_at,omitempty" time_format:"RFC3339"`
	UpdatedAt time.Time `json:"updated_at,omitempty" time_format:"RFC3339"`
}

type HTTPResponse struct {
	Msg string `json:"msg,omitempty" example:"any message"`
}

// type PostCustomerRequest struct {
// 	CustomerID string `uri:"customerid" binding:"required,uuid"`
// }

type PostCustomerResponse struct {
	ID string `json:"id" binding:"uuid"`
}

type Customer struct {
	Base `json:",inline"`
}

type PostPlaceResponse struct {
	ID string `json:"id" binding:"uuid"`
}

type Place struct {
	Base `json:",inline"`
}
