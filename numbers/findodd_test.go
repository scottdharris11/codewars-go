package numbers

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)
  
func DoTestFindOdd(a []int, exp int) {
	var ans = FindOdd(a)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("FindOdd", func() {
	It("FindOdd Test 1", func() {
		DoTestFindOdd( []int{20,1,-1,2,-2,3,3,5,5,1,2,4,20,4,-1,-2,5}, 5 )
	})
	It("FindOdd Test 2", func() {
		DoTestFindOdd( []int{1,1,2,-2,5,2,4,4,-1,-2,5}, -1 )
	})
	It("FindOdd Test 3", func() {
		DoTestFindOdd( []int{20,1,1,2,2,3,3,5,5,4,20,4,5}, 5 )
	})
	It("FindOdd Test 4", func() {
		DoTestFindOdd( []int{10}, 10 )
	})
	It("FindOdd Test 5", func() {
		DoTestFindOdd( []int{1,1,1,1,1,1,10,1,1,1,1}, 10 )
	})
	It("FindOdd Test 6", func() {
		DoTestFindOdd( []int{5,4,3,2,1,5,4,3,2,10,10}, 1 )
	})
})
