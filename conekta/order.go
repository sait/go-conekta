package conekta

type Order struct {
	ID              string          `json:"id,omitempty"`
	Object          string          `json:"object,omitempty"`
	CreatedAt       int64           `json:"created_at,omitempty"`
	UpdatedAt       int64           `json:"updated_at,omitempty"`
	Currency        string          `json:"currency,omitempty"`
	LineItems       []LineItem      `json:"line_items,omitempty"`
	ShippingLines   []ShippingLine  `json:"shipping_lines,omitempty"`
	TaxLines        []TaxLine       `json:"tax_lines,omitempty"`
	DiscountLines   []DiscountLine  `json:"discount_lines,omitempty"`
	PreAuthorize    bool            `json:"pre_authorize,omitempty"`
	CustomerInfo    Customer        `json:"customer_info,omitempty"`
	ShippingContact ShippingContact `json:"shipping_contact,omitempty"`
	Charges         []Charge        `json:"charges,omitempty"`
	Metadata        Metadata        `json:"metadata,omitempty"`
}

type LineItem struct {
	ID          string   `json:"id,omitempty"`
	Object      string   `json:"object,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	UnitPrice   float64  `json:"unit_price,omitempty"`
	Quantity    float64  `json:"quantity,omitempty"`
	Sku         string   `json:"sku,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Brand       string   `json:"brand,omitempty"`
}

type ShippingLine struct {
	Amunt          float64 `json:"amount,omitempty"`
	TrackingNumber string  `json:"tracking_number,omitempty"`
	Carrier        string  `json:"carrier,omitempty"`
	Method         string  `json:"method,omitempty"`
}

type TaxLine struct {
	Description string  `json:"description,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
}

type DiscountLine struct {
	Code   string  `json:"code,omitempty"`
	Type   string  `json:"type,omitempty"`
	Amount float64 `json:"amount,omitempty"`
}

type Metadata map[string]string

type Charge struct {
	PaymentMethod PaymentMethod `json:"payment_method,omitempty"`
}

type PaymentMethod struct {
	Type string `json:"type,omitempty"`
}

func (o *Order) Create() (statusCode int) {
	return request("POST", "/orders", o)
}
