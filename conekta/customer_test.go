package conekta_test

import (
	"testing"
	"os"

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
	conekta.ApiKey = os.Getenv("CONEKTAKEY")
	Context("Post customer", func() {
		It("Should response 200", func() {
			//New customer
			customer := new(conekta.Customer)
			customer.Name = "Fulanito Pérez"
			customer.Email = "fulanito@conekta.com"
			customer.Phone = "+52181818181"
			//Testing payment src
			payment := conekta.PaymentSource{
				Type:    "card",
				TokenID: "tok_test_visa_4242",
			}
			customer.PaymentSources = append(customer.PaymentSources, payment)
			//Send to conekta
			statusCode, _, _ := customer.Create()
			Expect(statusCode).Should(Equal(200))
		})
	})
})
