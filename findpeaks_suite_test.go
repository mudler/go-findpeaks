package findpeaks_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFindPeaks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Findpeaks test suite")
}
