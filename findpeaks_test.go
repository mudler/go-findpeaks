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
			Expect(peaks).To(Equal([]int{1}))
		})

		It("describe the sequence", func() {
			testCase := []LowAndHigh{
				{Value: 1.1, Type: H},
				{Value: 1.0, Type: L},
				{Value: 1.9, Type: H},
				{Value: 0.4, Type: S},
				{Value: 0.4, Type: S},
				{Value: 0.4, Type: S},
				{Value: 0.6, Type: S},
				{Value: 0.6, Type: S},
				{Value: 0.8, Type: H},
			}
			peaks := LowAndHighs(Sequence(testCase).Float64())
			Expect(peaks).To(Equal(testCase))
		})
	})
})
