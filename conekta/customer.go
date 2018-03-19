package conekta

type Customer struct {
	Name             string            `json:"name"`
	Email            string            `json:"email"`
	Phone            string            `json:"phone,omitempty"`
	PaymentSources   []PaymentSource   `json:"payment_sources,omitempty"`
	ShippingContacts []ShippingContact `json:"shipping_contacts,omitempty"`
}

type PaymentSource struct {
	TokenID string `json:"token_id"`
	Type    string `json:"type"`
}

type ShippingContact struct {
	Phone    string  `json:"phone"`
	Receiver string  `json:"receiver,omitempty"`
	Address  Address `json:"address"`
}

type Address struct {
	Street1    string `json:"street1"`
	Street2    string `json:"street1,omitempty"`
	City       string `json:"city,omitempty"`
	State      string `json:"string,omitempty"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}

func (c *Customer) Post() (statusCode int) {
	return request("POST", conektaUrl+"/customers", c)
}
