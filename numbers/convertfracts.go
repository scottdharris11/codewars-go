package numbers

import "fmt"

func ConvertFracts(a [][]int) string {
	// Find largest denominator after simplification
	lDenom := 0
	for _, r := range a {
		gcd := GCD( r[0], r[1] )
		if ( gcd != r[1] ) {
			r[0] = r[0] / gcd
			r[1] = r[1] / gcd
		}
		if ( r[1] > lDenom ) {
			lDenom = r[1]
		}
	}

	// Look for smallest multiple of the largest denom that is 
	// evenly divisible by each denomiator
	denom := lDenom
	allEven := false
	for i := 1; ! allEven; i++ {
		denom = lDenom * i
		allEven = true
		for _, r := range a {
			if ( ( denom % r[1] ) != 0 ) {
				allEven = false
				break
			}
		}
	}

	// Build out
	output := ""
	num := 0
	for _, r := range a {
		num = (denom/r[1]) * r[0]
		output += fmt.Sprintf( "(%d,%d)", num, denom )
	}
    return output
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}