package main

func MaximumSubarraySum(numbers []int) int {
	max := 0
	subMax := 0
	for _, num := range numbers {
		subMax += num
		if ( max < subMax ) {
			max = subMax
		}
		if ( subMax < 0 ) {
			subMax = 0
		}
	}
	return max
}