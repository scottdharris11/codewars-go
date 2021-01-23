package main

import "math"

func SquaresInRect(lng int, wdth int) []int {
	if ( lng == wdth ) {
		return nil
	}
	squares := []int {}
	wrkMaxLng := float64( lng )
	wrkMaxWdth := float64( wdth )
	for {
		sq := math.Floor( math.Sqrt( float64( wrkMaxLng * wrkMaxWdth ) ) )
		if ( sq >= wrkMaxLng ) {
			sq = wrkMaxLng
			wrkMaxWdth -= sq
		} else if ( sq >= wrkMaxWdth ) {
			sq = wrkMaxWdth
			wrkMaxLng -= sq
		}
		if ( sq > 0 ) {
			squares = append( squares, int( sq ) )
		} else {
			return squares
		}
	}
}