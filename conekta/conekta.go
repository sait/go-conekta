package conekta

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
	"encoding/json"
)

type Token struct {
	Id       string `json:"id"`
	Object   string `json:"object"`
	Used     bool   `json:"used"`
	LiveMode bool   `json:"livemode"`
}

type Error struct {
	Type             string `json:"id"`
	Message          string `json:"message"`
	MessagePurchaser string `json:"message_to_purchaser"`
	ErrorCode        string `json:"error_code"`
	Param            string `json:"param"`
}

var (
	ApiKey, ApiVersion = "", "2.0.0"
)

const (
	conektaUrl = "https://api.conekta.io"
)

func request(method, url string, v interface{}) {
	jsonPayload, err := json.Marshal(v)
	if err != nil {
		return
	}
	payload := bytes.NewReader(jsonPayload)
	req, _ := http.NewRequest(method, url, payload)
	req.Header.Add("accept", "application/vnd.conekta-v"+ApiVersion+"+json")
	req.SetBasicAuth(ApiKey, "")
	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(res)
	fmt.Println(string(body))
}
