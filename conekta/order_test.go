package conekta_test

import (
	"os"
	"testing"

	"github.com/sait/go-conekta/conekta"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOrder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handling order")
}

var _ = Describe("Handle order", func() {
	//Testing key
	conekta.ApiKey = os.Getenv("CONEKTAKEY")
	Context("Create order to a customer", func() {
		It("Should response 200", func() {
			//New Order
			order := new(conekta.Order)
			item := conekta.LineItem{
				Name:      "Tortugas",
				UnitPrice: 1000,
				Quantity:  12,
			}
			order.LineItems = append(order.LineItems, item)
			shipping := conekta.ShippingLine{
				Amount:  1500,
				Carrier: "FEDEX",
			}
			order.ShippingLines = append(order.ShippingLines, shipping)
			order.Currency = "MXN"
			//testing customer id
			order.CustomerInfo.CustomerID = "cus_2iFzZsBLvnx11gyXy"
			order.ShippingContact = &conekta.ShippingContact{
				Address: conekta.Address{
					Street1:    "Calle 123, int 2",
					PostalCode: "06100",
					Country:    "MX",
				},
			}
			//Adding some metadata
			order.Metadata = conekta.Metadata{
				"reference": "12987324097",
				"more_info": "lalalalala",
				"hello":     "world",
			}
			charge := conekta.Charge{
				PaymentMethod: conekta.PaymentMethod{
					Type: "oxxo_cash",
				},
			}
			order.Charges = append(order.Charges, charge)
			statusCode, _ := order.Create()
			Expect(statusCode).Should(Equal(200))
		})
	})
	Context("Create a order directly", func() {
		It("Should response 200", func() {
			//New Order
			order := new(conekta.Order)
			item := conekta.LineItem{
				Name:        "Churros Locos",
				Description: "Made in Mexico.",
				UnitPrice:   20000,
				Quantity:    2,
			}
			order.LineItems = append(order.LineItems, item)
			order.Currency = "MXN"
			order.Metadata = conekta.Metadata{
				"test": "extra_info",
				"hola": "mundo",
			}
			charge := conekta.Charge{
				PaymentMethod: conekta.PaymentMethod{
					Type:    "card",
					TokenId: "tok_test_visa_4242",
				},
			}
			order.Charges = append(order.Charges, charge)
			order.CustomerInfo.Name = "Fulanito Pérez"
			order.CustomerInfo.Email = "fulanito@conekta.com"
			order.CustomerInfo.Phone = "+52181818181"
			statusCode, _ := order.Create()
			Expect(statusCode).Should(Equal(200))
		})
	})
	Context("Update order", func() {
		It("Should response 200", func() {
			order := new(conekta.Order)
			order.ID = "ord_2iGHf3etsiNEji8Ry"
			order.Currency = "MXN"
			statusCode, _ := order.Update()
			Expect(statusCode).Should(Equal(200))
		})
	})
	Context("Capture order", func() {
		It("Should response 200", func() {
			order := new(conekta.Order)
			order.ID = "ord_2iGPs5fX4uTnqhCJX"
			statusCode, _ := order.Capture()
			//A preauthorized order can captured only once
			Expect(statusCode).Should(Equal(428))
		})
	})
	Context("Refound order", func() {
		It("Should response 200", func() {
			order := new(conekta.Order)
			order.ID = "ord_2iY6n3yfqKWXdBJ4n"
			order.Reason = "requested_by_client"
			order.Amunt = 100
			statusCode, _ := order.Refund()
			Expect(statusCode).Should(Equal(200))
		})
	})
})
