package findpeaks_test

import (
	. "github.com/mudler/go-findpeaks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Profits", func() {
	Context("evaluating sequences", func() {
		It("find proficts", func() {
			Expect(MaxProfit([]float64{1.1, 1.0, 1.9, 0.4, 0.4, 0.4, 0.6, 0.6, 0.8}, 1.0)).To(Equal(2.8))
		})
	})
})
