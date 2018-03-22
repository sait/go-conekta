package conekta

import (
	"bytes"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"fmt"
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

var (
	ApiKey, ApiVersion = "", "2.0.0"
)

const (
	conektaUrl = "https://api.conekta.io"
)

func request(method, path string, v interface{}) (statusCode int, response []byte) {
	jsonPayload, err := json.MarshalIndent(v, "", "   ")
	if err != nil {
		return
	}
	fmt.Printf("Este es el JSON que le mando==================\n%s\n", jsonPayload)
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
		panic(err)
	}
}