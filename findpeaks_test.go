package findpeaks_test

import (
	. "github.com/mudler/go-findpeaks"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FindPeaks", func() {
	Context("evaluating sequences", func() {
		It("find peaks in a sequence", func() {
			peaks := FindPeaks([]float64{1.1, 1.0, 1.9, 0.4})
			Expect(peaks).To(Equal([]int{0, 2}))
		})
		It("find low points in a sequence", func() {
			peaks := FindPeaks(ReverseSequence([]float64{1.1, 1.0, 1.9, 0.4, 0.4, 0.4, 0.6, 0.6, 0.8}))
			Expect(peaks).To(Equal([]int{1, 3, 4, 5, 7}))
		})
	})
})
