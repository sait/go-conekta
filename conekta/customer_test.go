package conekta_test

import (
	"testing"

	"github.com/sait/go-conekta/conekta"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCustomer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Create a customer")
}

var _ = Describe("Creating customer", func() {
	//Testing key
	conekta.ApiKey = "key_eYvWV7gSDkNYXsmr"
	Context("Post customer", func() {
		It("Should response 200", func() {
			//New customer
			customer := new(conekta.Customer)
			customer.Name = "Fulanito Pérez"
			customer.Email = "fulanito@conekta.com"
			customer.Phone = "+52181818181"
			//Testing payment
			payment := conekta.PaymentSource{
				Type:    "card",
				TokenID: "tok_test_visa_4242",
			}
			customer.PaymentSources = append(customer.PaymentSources, payment)
			statusCode := customer.Post()
			Expect(statusCode).Should(Equal(200))
		})
	})
})
