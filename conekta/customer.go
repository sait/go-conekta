package conekta

import (
	"encoding/json"
)

type Customer struct {
	CustomerID       string            `json:"customer_id,omitempty"`
	Name             string            `json:"name,omitempty"`
	Phone            string            `json:"phone,omitempty"`
	Email            string            `json:"email,omitempty"`
	Corporate        bool              `json:"corporate,omitempty"`
	PaymentSources   []PaymentSource   `json:"payment_sources,omitempty"`
	ShippingContacts []ShippingContact `json:"shipping_contacts,omitempty"`
}

type PaymentSource struct {
	TokenID string `json:"token_id,omitempty"`
	Type    string `json:"type,omitempty"`
}

type ShippingContact struct {
	ID             string   `json:"id,omitempty"`
	Object         string   `json:"object,omitempty"`
	CreatedAt      int64    `json:"created_at,omitempty"`
	UpdatedAt      int64    `json:"updated_at,omitempty"`
	Phone          string   `json:"phone,omitempty"`
	Receiver       string   `json:"receiver,omitempty"`
	BetweenStreets string   `json:"between_streets,omitempty"`
	Address        Address  `json:"address,omitempty"`
	Metadata       Metadata `json:"metadata,omitempty"`
}

type Address struct {
	Street1     string `json:"street1,omitempty"`
	Street2     string `json:"street2,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"string,omitempty"`
	Country     string `json:"country,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	Residential bool   `json:"residential,omitempty"`
	Object      string `json:"object,omitempty"`
}

func (c *Customer) Create() (statusCode int, conektaError ConektaError, conektaResponse ConektaResponse) {
	statusCode, response := request("POST", "/customers", c)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &conektaResponse)
		checkError(err)
	}
	return
}
