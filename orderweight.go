package main

import (
	"strings"
	"strconv"
	"sort"
)

func OrderWeight(strng string) string {
	weightStrs := strings.Split( strng, " " )
	sort.SliceStable(weightStrs, func(i, j int) bool {
		weightI := WeightForString( weightStrs[i] )
		weightJ := WeightForString( weightStrs[j] )
		return weightI < weightJ || ( weightI == weightJ && weightStrs[i] < weightStrs[j] )
	})
	return strings.Join( weightStrs, " " ) 
}

func WeightForString( s string ) int {
	weight := 0
	for _, c := range s {
		charWeight, _ := strconv.Atoi( string( c ) )
		weight += charWeight
	}
	return weight
}
