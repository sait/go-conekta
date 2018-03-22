package conekta

import (
	"encoding/json"
)

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
	Livemode        bool            `json:"livemode,omitempty"`
	PreAuthorize    bool            `json:"pre_authorize,omitempty"`
	ShippingContact ShippingContact `json:"-"`
	Amunt           float64         `json:"amount,omitempty"`
	Reason          string          `json:"reason,omitempty"`
	AmountRefunded  float64         `json:"amount_refunded,omitempty"`
	PaymentStatus   string          `json:"payment_status,omitempty"`
	CustomerInfo    Customer        `json:"customer_info,omitempty"`
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
	Tags        Tags     `json:"tags,omitempty"`
	Brand       string   `json:"brand,omitempty"`
	ParentID    string   `json:"parent_id,omitempty"`
	Metadata    Metadata `json:"metadata,omitempty"`
}

type ShippingLine struct {
	ID             string   `json:"id,omitempty"`
	Object         string   `json:"object,omitempty"`
	Amunt          float64  `json:"amount,omitempty"`
	TrackingNumber string   `json:"tracking_number,omitempty"`
	Carrier        string   `json:"carrier,omitempty"`
	Method         string   `json:"method,omitempty"`
	ParentID       string   `json:"parent_id,omitempty"`
	Metadata       Metadata `json:"metadata,omitempty"`
}

type TaxLine struct {
	ID          string   `json:"id,omitempty"`
	Object      string   `json:"object,omitempty"`
	Description string   `json:"description,omitempty"`
	Amount      float64  `json:"amount,omitempty"`
	ParentID    string   `json:"parent_id,omitempty"`
	Metadata    Metadata `json:"metadata,omitempty"`
}

type DiscountLine struct {
	ID       string   `json:"id,omitempty"`
	Object   string   `json:"object,omitempty"`
	Code     string   `json:"code,omitempty"`
	Type     string   `json:"type,omitempty"`
	Amount   float64  `json:"amount,omitempty"`
	ParentID string   `json:"parent_id,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"`
}

type Tags map[string]string

type Metadata map[string]string

type Charge struct {
	ID                  string        `json:"id,omitempty"`
	Object              string        `json:"object,omitempty"`
	CreatedAt           int64         `json:"created_at,omitempty"`
	UpdatedAt           int64         `json:"updated_at,omitempty"`
	ExpiresAt           int64         `json:"expires_at,omitempty"`
	Currency            string        `json:"currency,omitempty"`
	Amount              float64       `json:"amount,omitempty"`
	MonthlyInstallments float64       `json:"monthly_installments,omitempty"`
	Livemode            bool          `json:"livemode,omitempty"`
	Status              string        `json:"status,omitempty"`
	Fee                 float64       `json:"fee,omitempty"`
	OrderID             string        `json:"order_id,omitempty"`
	PaymentMethod       PaymentMethod `json:"payment_method,omitempty"`
}

type PaymentMethod struct {
	Type    string `json:"type,omitempty"`
	TokenId string `json:"token_id,omitempty"`
}

// Creates a new Order
func (o *Order) Create() (statusCode int, conektaError *ConektaError) {
	statusCode, response := request("POST", "/orders", o)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	}
	return
}

// Updates an existing Order
func (o *Order) Update() (statusCode int, conektaError *ConektaError) {
	statusCode, response := request("PUT", "/orders/"+o.ID, o)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	}
	return
}

// Process a pre-authorized order.
func (o *Order) Capture() (statusCode int, conektaError *ConektaError) {
	statusCode, response := request("POST", "/orders/"+o.ID+"/capture", nil)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	}
	return
}

// A Refund details the amount and reason why an order was refunded.
func (o *Order) Refund() (statusCode int, conektaError *ConektaError) {
	statusCode, response := request("POST", "/orders/"+o.ID+"/refunds", o)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	}
	return
}
