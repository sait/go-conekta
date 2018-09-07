package conekta_test

import (
	"os"
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
	conekta.ApiKey = os.Getenv("CONEKTAKEY")
	var cusid string
	Context("Post customer", func() {
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
		statusCode, _, response := customer.Create()
		cusid = response.ID
		It("Should response 200", func() {
			Expect(statusCode).Should(Equal(200))
		})
	})

	Context("Update a customer", func() {
		//New customer
		customer := new(conekta.Customer)
		customer.CustomerID = cusid
		customer.Name = "Zutano Pérez"
		customer.Email = "zutano@conekta.com"
		customer.Phone = "+52181818181"
		statusCode, _, _ := customer.Update()
		It("Should response 200", func() {
			Expect(statusCode).Should(Equal(200))
		})
	})

	Context("Create a subscription", func() {
		customer := new(conekta.Customer)
		customer.CustomerID = cusid
		statusCode, _, _ := customer.CreateSubscription("399")
		It("Should response 200", func() {
			Expect(statusCode).Should(Equal(200))
		})
	})

	Context("Update a subscription", func() {
		customer := new(conekta.Customer)
		customer.CustomerID = cusid
		statusCode, _, _ := customer.UpdateSubscription("400")
		It("Should response 200", func() {
			Expect(statusCode).Should(Equal(200))
		})
	})

	Context("Pause a subscription", func() {
		customer := new(conekta.Customer)
		customer.CustomerID = cusid
		statusCode, _, _ := customer.PauseSubscription()
		It("Should response 200", func() {
			Expect(statusCode).Should(Equal(200))
		})
	})

	Context("Resume a subscription", func() {
		customer := new(conekta.Customer)
		customer.CustomerID = cusid
		statusCode, _, _ := customer.ResumeSubscription()
		It("Should response 200", func() {
			Expect(statusCode).Should(Equal(200))
		})
	})

	Context("Cancel a subscription", func() {
		customer := new(conekta.Customer)
		customer.CustomerID = cusid
		statusCode, _, _ := customer.CancelSubscription()
		It("Should response 200", func() {
			Expect(statusCode).Should(Equal(200))
		})
	})

	Context("Delete a customer", func() {
		customer := new(conekta.Customer)
		customer.CustomerID = cusid
		statusCode, _, _ := customer.Delete()
		It("Should response 200", func() {
			Expect(statusCode).Should(Equal(200))
		})
	})
})
