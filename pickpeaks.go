package main

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
	peaks := PosPeaks{}
	if ( len( array ) > 0 ) {
		lastNum := array[0]
		lastPos := 0
		asc := false
		for idx, num := range array {
			if ( num > lastNum ) {
				lastNum = num
				lastPos = idx
				asc = true
			} else if ( num < lastNum ) {
				if ( asc ) {
					peaks.Peaks = append( peaks.Peaks, lastNum )
					peaks.Pos = append( peaks.Pos, lastPos )
					asc = false
				}
				lastNum = num
				lastPos = idx
			}
		}
	}
	if ( peaks.Peaks == nil ) {
		peaks.Peaks = []int{}
		peaks.Pos = []int{}
	}
	return peaks
}