package main

import (
	"fmt"
	"slices"
)

// isIncreasing reports whether nums is strictly increasing (each element < next).
func isIncreasing(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			return false
		}
	}
	return true
}

// isDecreasing reports whether nums is strictly decreasing (each element > next).
func isDecreasing(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] <= nums[i+1] {
			return false
		}
	}
	return true
}

// isTrionic reports whether nums has shape: increasing [0..p], decreasing [p..q], increasing [q..end].
// p and q are chosen from values in nums that lie in (1, len-1); q = max of those, p = min.
func isTrionic(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	var candidates []int
	for _, v := range nums {
		if v > 1 && v < n-1 {
			candidates = append(candidates, v)
		}
	}
	if len(candidates) == 0 {
		return false
	}

	p, q := slices.Min(candidates), slices.Max(candidates)
	if p >= q {
		return false
	}

	// Segment 1: [0..p] increasing; segment 2: [p..q] decreasing; segment 3: [q..n) increasing.
	return isIncreasing(nums[0:p+1]) &&
		isDecreasing(nums[p:q+1]) &&
		isIncreasing(nums[q:n])
}

func main() {
	fmt.Println(isTrionic([]int{2, 1, 3}))
}
