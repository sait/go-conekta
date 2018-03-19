package conekta

type Order struct {
	LineItems       []LineItem      `json:"line_items,omitempty"`
	ShippingLines   []ShippingLine  `json:"shipping_lines,omitempty"`
	Currency        string          `json:"currency,omitempty"`
	CustomerInfo    CustomerInfo    `json:"customer_info,omitempty"`
	ShippingContact ShippingContact `json:"shipping_contact,omitempty"`
	Metadata        Metadata        `json:"metadata,omitempty"`
	Charges         []Charge        `json:"charges,omitempty"`
}

type LineItem struct {
	Name      string  `json:"name,omitempty"`
	UnitPrice float64 `json:"unit_price,omitempty"`
	Quantity  float64 `json:"quantity,omitempty"`
}

type ShippingLine struct {
	Amunt   float64 `json:"amount,omitempty"`
	Carrier string  `json:"carrier,omitempty"`
}

type CustomerInfo struct {
	CustomerID string `json:"customer_id,omitempty"`
}

type Metadata struct {
	Reference string `json:"reference,omitempty"`
	Moreinfo  string `json:"more_info,omitempty"`
}

type Charge struct {
	PaymentMethod PaymentMethod `json:"payment_method,omitempty"`
}

type PaymentMethod struct {
	Type string `json:"type,omitempty"`
}

func (o *Order) Post() (statusCode int) {
	return request("POST", conektaUrl+"/orders", o)
}
