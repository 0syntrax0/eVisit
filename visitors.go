package main

import (
	"sort"
)

// checks if a given IP is a top visitor
func updateTopVisitors(ip string, count int) {
	// - lock the mutex!

	// - sort 'topIpVisitors' slice in DESC order
	// sort desc
	sort.SliceStable(topIpVisitors, func(i, j int) bool {
		return topIpVisitors[i][j] < topIpVisitors[j][j]
	})

	// - check if current count is greater than the minimum count: len(topIpVisitors)-1
	// // - get slice key of ip with closest count to it add it to the slice above
	// // - if slice length is above "topVisitorCount", drop the bottom slices

	// - release the mutex
}
