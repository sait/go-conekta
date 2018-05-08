package conekta_test

import (
	"testing"
	"github.com/sait/go-conekta/conekta"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConekta(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Format amount")
}

var _ = Describe("Formatting amount", func() {
	Context("Formatting some numbers", func(){
		var testnum1 float64
		testnum1 = 352
		formatted1, _ := conekta.ConektaFormatAmount(testnum1)
		It("Should be formatted", func(){
			var equal1 int64
			equal1 = 35200
			Expect(formatted1).Should(Equal(equal1))
		})

		var testnum2 float64
		testnum2 = 150.50
		formatted2, _ := conekta.ConektaFormatAmount(testnum2)
		It("Should be formatted", func(){
			var equal2 int64
			equal2 = 15050
			Expect(formatted2).Should(Equal(equal2))
		})
	})
})
