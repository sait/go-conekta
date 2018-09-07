package conekta

import (
	"encoding/json"
)

type Customer struct {
	CustomerID               string            `json:"customer_id,omitempty"`
	Name                     string            `json:"name,omitempty"`
	Phone                    string            `json:"phone,omitempty"`
	Email                    string            `json:"email,omitempty"`
	Corporate                bool              `json:"corporate,omitempty"`
	DefaultPaymentSourceID   string            `json:"default_payment_source_id,omitempty"`
	DefaultShippingContactID string            `json:"default_shipping_contact_id,omitempty"`
	PaymentSources           []PaymentSource   `json:"payment_sources,omitempty"`
	ShippingContacts         []ShippingContact `json:"shipping_contacts,omitempty"`
}

type PaymentSource struct {
	ID        string `json:"id,omitempty"`
	Object    string `json:"object,omitempty"`
	TokenID   string `json:"token_id,omitempty"`
	Type      string `json:"type,omitempty"`
	CreatedAt int64  `json:"created_at,omitempty"`
	Last4     string `json:"last4,omitempty"`
	Name      string `json:"name,omitempty"`
	ExpMonth  string `json:"exp_month,omitempty"`
	ExpYear   string `json:"exp_year,omitempty"`
	Brand     string `json:"brand,omitempty"`
	ParentID  string `json:"parent_id,omitempty"`
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

type Subscription struct {
	ID                string `json:"id,omitempty"`
	Object            string `json:"object,omitempty"`
	CreatedAt         int64  `json:"created_at,omitempty"`
	UpdatedAt         int64  `json:"updated_at,omitempty"`
	PausedAt          int64  `json:"paused_at,omitempty"`
	BillingCycleStart int64  `json:"billing_cycle_start,omitempty"`
	BillingCycleEnd   int64  `json:"billing_cycle_end,omitempty"`
	TrialStart        int64  `json:"trial_start,omitempty"`
	TrialEnd          int64  `json:"trial_end,omitempty"`
	PlanID            string `json:"plan_id,omitempty"`
	Status            string `json:"status,omitempty"`
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

func (c *Customer) Update() (statusCode int, conektaError ConektaError, conektaResponse ConektaResponse) {
	statusCode, response := request("PUT", "/customers/"+c.CustomerID, c)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &conektaResponse)
		checkError(err)
	}
	return
}

func (c *Customer) Delete() (statusCode int, conektaError ConektaError, conektaResponse ConektaResponse) {
	statusCode, response := request("DELETE", "/customers/"+c.CustomerID, nil)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &conektaResponse)
		checkError(err)
	}
	return
}

func (c *Customer) CreateSubscription(plan string) (statusCode int, conektaError ConektaError, subscription Subscription) {
	statusCode, response := request("POST", "/customers/"+c.CustomerID+"/subscription", body{"plan": plan})
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &subscription)
		checkError(err)
	}
	return
}

func (c *Customer) UpdateSubscription(plan string) (statusCode int, conektaError ConektaError, subscription Subscription) {
	statusCode, response := request("PUT", "/customers/"+c.CustomerID+"/subscription", body{"plan": plan})
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &subscription)
		checkError(err)
	}
	return
}

func (c *Customer) PauseSubscription() (statusCode int, conektaError ConektaError, subscription Subscription) {
	statusCode, response := request("POST", "/customers/"+c.CustomerID+"/subscription/pause", nil)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &subscription)
		checkError(err)
	}
	return
}

func (c *Customer) ResumeSubscription() (statusCode int, conektaError ConektaError, subscription Subscription) {
	statusCode, response := request("POST", "/customers/"+c.CustomerID+"/subscription/resume", nil)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &subscription)
		checkError(err)
	}
	return
}

func (c *Customer) CancelSubscription() (statusCode int, conektaError ConektaError, subscription Subscription) {
	statusCode, response := request("POST", "/customers/"+c.CustomerID+"/subscription/cancel", nil)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &subscription)
		checkError(err)
	}
	return
}

func (c *Customer) CreatePaymentSource(paymentSource PaymentSource) (statusCode int, conektaError ConektaError, paymentResponse PaymentSource) {
	statusCode, response := request("POST", "/customers/"+c.CustomerID+"/payment_sources/", paymentSource)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &paymentResponse)
		checkError(err)
	}
	return
}

func (c *Customer) DeletePaymentSource(paymentSourceID string) (statusCode int, conektaError ConektaError, paymentResponse PaymentSource) {
	statusCode, response := request("DELETE", "/customers/"+c.CustomerID+"/payment_sources/"+paymentSourceID, nil)
	if statusCode != 200 {
		err := json.Unmarshal(response, &conektaError)
		checkError(err)
	} else {
		err := json.Unmarshal(response, &paymentResponse)
		checkError(err)
	}
	return
}
