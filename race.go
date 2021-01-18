package main

import (
	"math"
)

func Race(v1, v2, g int) [3]int {
	result := [3]int{-1, -1, -1}
	if ( v1 < v2 ) {
		secsInHour := 60*60
		feetPerSec := float64( v2 - v1 ) / float64( secsInHour )
		secsToCatch := int( float64( g ) / float64( feetPerSec ) )

		hoursToCatch := int( math.Floor( float64( secsToCatch ) / float64( secsInHour ) ) )
		secsToCatch = secsToCatch - ( hoursToCatch * secsInHour )

		minsToCatch := int( math.Floor( float64( secsToCatch ) / 60.0 ) )
		secsToCatch = secsToCatch - ( minsToCatch * 60 )

		result[0] = hoursToCatch
		result[1] = minsToCatch
		result[2] = secsToCatch
	}
	return result
}