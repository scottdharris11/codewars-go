package numbers

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)
  
func DoTestConvertFracts(a [][]int, exp string) {
	var ans = ConvertFracts(a)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("ConvertFracts", func() {
	It("Basic test 1", func() {
		var lst = [][]int{{1, 2}, {1, 3}, {1, 4}}
		DoTestConvertFracts(lst, "(6,12)(4,12)(3,12)")
	})
	It("Basic test 2", func() {
		var lst = [][]int{{69, 130}, {87, 1310}, {30, 40}}
		DoTestConvertFracts(lst, "(18078,34060)(2262,34060)(25545,34060)")
	})
})