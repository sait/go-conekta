package conekta

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
	Phone    string  `json:"phone,omitempty"`
	Receiver string  `json:"receiver,omitempty"`
	Address  Address `json:"address,omitempty"`
}

type Address struct {
	Street1    string `json:"street1,omitempty"`
	Street2    string `json:"street2,omitempty"`
	City       string `json:"city,omitempty"`
	State      string `json:"string,omitempty"`
	Country    string `json:"country,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
}

func (c *Customer) Create() (statusCode int) {
	return request("POST", "/customers", c)
}
