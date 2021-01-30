package numbers

import "sort"

func FindOdd(seq []int) int {
	sort.Ints( seq )
	previous := 0
	count := 0
	for idx, num := range seq {
		if ( idx > 0 && previous != num ) {
			if ( count % 2 == 1 ) {
				return previous
			}
			count = 0
		}
		count++
		previous = num
	}
    return previous
}