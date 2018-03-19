package conekta

//TODO: structura de la orden
type Order struct {
	Data string
}

func (o *Order) Post() (statusCode int) {
	return request("POST", conektaUrl+"/orders", o)
}
