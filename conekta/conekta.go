package conekta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type ConektaError struct {
	Object  string   `json:"object,omitempty"`
	Type    string   `json:"type,omitempty"`
	LogId   string   `json:"log_id,omitempty"`
	Details []Detail `json:"details,omitempty"`
}

type Detail struct {
	Debug_message string `json:"debug_message,omitempty"`
	Message       string `json:"message,omitempty"`
	Code          string `json:"code,omitempty"`
}

type body map[string]interface{}

type ConektaResponse struct {
	Livemode        *bool           `json:"livemode,omitempty"`
	Amount          int64           `json:"amount,omitempty"`
	Currency        string          `json:"currency,omitempty"`
	PaymentStatus   string          `json:"payment_status,omitempty"`
	AmountRefunded  int64           `json:"amount_refunded,omitempty"`
	CustomerInfo    CustomerInfo    `json:"customer_info,omitempty"`
	ShippingContact ShippingContact `json:"shipping_contact,omitempty"`
	Object          string          `json:"object,omitempty"`
	ID              string          `json:"id,omitempty"`
	Metadata        Metadata        `json:"metadata,omitempty,omitempty"`
	CreatedAt       int64           `json:"created_at,omitempty"`
	UpdatedAt       int64           `json:"updated_at,omitempty"`
	LineItems       LineItems       `json:"line_items"`
	ShippingLines   ShippingLines   `json:"shipping_lines,omitempty"`
	Charges         Charges         `json:"charges,omitempty"`
}

type Charges struct {
	Object  string   `json:"object,omitempty"`
	HasMore *bool    `json:"has_more,omitempty"`
	Total   int64    `json:"total,omitempty"`
	Data    []Charge `json:"data,omitempty"`
}

type ShippingLines struct {
	Object  string         `json:"object,omitempty"`
	HasMore *bool          `json:"has_more,omitempty"`
	Total   int64          `json:"total,omitempty"`
	Data    []ShippingLine `json:"data,omitempty"`
}

type CustomerInfo struct {
	Email      string `json:"email,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Name       string `json:"name,omitempty"`
	Corporate  *bool  `json:"corporate,omitempty"`
	CustomerID string `json:"customer_id,omitempty"`
	Object     string `json:"object,omitempty"`
}

type LineItems struct {
	Object  string     `json:"object,omitempty"`
	HasMore *bool      `json:"has_more,omitempty"`
	Total   int64      `json:"total,omitempty"`
	Data    []LineItem `json:"data,omitempty"`
}

type AntifraudInfo map[string]string

var (
	ApiKey, ApiVersion = "", "2.0.0"
)

const (
	conektaUrl = "https://api.conekta.io"
)

func request(method, path string, v interface{}) (statusCode int, response []byte) {
	jsonPayload, err := json.Marshal(v)
	if err != nil {
		return
	}
	payload := bytes.NewReader(jsonPayload)
	req, _ := http.NewRequest(method, conektaUrl+path, payload)
	req.Header.Add("accept", "application/vnd.conekta-v"+ApiVersion+"+json")
	req.SetBasicAuth(ApiKey, "")
	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return res.StatusCode, body
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("There's an error in Conekta Wrapper: %v\n", err)
	}
}

func ConektaFormatAmount(value float64) (formatted int64, err error) {
	strnum := fmt.Sprintf("%.2f", value)
	strnum = strings.Replace(strnum, ".", "", -1)
	formatted, err = strconv.ParseInt(strnum, 10, 64)
	return
}

func ConektaFormatToFloat64(conektaFormatted int64) float64 {
	return float64(conektaFormatted) / 100
}