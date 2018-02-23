package models

import "time"

//
// This is used for returning a response with a single order as body
//
// swagger:response pingResponse
type ResPingBody struct {
	// PingBody
	//
	// in:body
	Body RspPing
}

// swagger:response pingResponseArr
type ResPingBodyArr struct {
	// PingBody
	//
	// in:body
	Body []RspPing
}

type RspPing struct {
	// Return Message
	Msg string `json:"msgstr"`

	// The time of Current Request
	CurrentTime time.Time `json:"current_time"`

	//
	// in: body
	Payload *PintModel `json:"order"`
}

// An Order for one or more products by a user.
// swagger:model order
type PintModel struct {
	// ID of the order
	//
	// required: true
	ID int64 `json:"id"`

	// the id of the user who placed the order.
	//
	// required: true
	UserID int64 `json:"user_id"`

	// items for this order
	// mininum items: 1
	OrderItems []struct {

		// the id of the product to order
		//
		// required: true
		ProductID int64 `json:"product_id"`

		// the quantity of this product to order
		//
		// required: true
		// minimum: 1
		Quantity int32 `json:"qty"`
	} `json:"items"`
}
