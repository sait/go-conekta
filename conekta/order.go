package conekta

//TODO: structura de la orden
type Order struct {
	Data string
}

func (o *Order) Post() {
	request("POST", conektaUrl+"/orders", o)
}
