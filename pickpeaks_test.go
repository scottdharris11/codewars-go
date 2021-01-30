package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func DoTestPickPeaks(array []int, posPeaks PosPeaks) {
	var ans = PickPeaks(array)
	Expect(ans).To(Equal(posPeaks))
}

var _ = Describe("PickPeaks Tests", func() {
		It("should support finding peaks", func() {
			DoTestPickPeaks(
				[]int{1, 2, 3, 6, 4, 1, 2, 3, 2, 1},
				PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}} )
		})
		
		It("should support finding peaks, but should ignore peaks on the edge of the array", func() {
			DoTestPickPeaks(
				[]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3},
				PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}},
			)
		})
	
		It("should support finding peaks; if the peak is a plateau, it should only return the position of the first element of the plateau", func() {
			DoTestPickPeaks(
				[]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1},
				PosPeaks{Pos: []int{3, 7, 10}, Peaks: []int{6, 3, 2}},
			)
		})
	
		It("should support finding peaks; if the peak is a plateau, it should only return the position of the first element of the plateau", func() {
			DoTestPickPeaks(
				[]int{2, 1, 3, 1, 2, 2, 2, 2, 1},
				PosPeaks{Pos: []int{2, 4}, Peaks: []int{3, 2}},
			)
		})
	
		It("should support finding peaks, but should ignore peaks on the edge of the array", func() {
			DoTestPickPeaks(
				[]int{2, 1, 3, 1, 2, 2, 2, 2},
				PosPeaks{Pos: []int{2}, Peaks: []int{3}},
			)
		})

		It("should handle empty input", func() {
			DoTestPickPeaks(
				[]int{},
				PosPeaks{Pos: []int{}, Peaks: []int{}},
			)
		})
	})

