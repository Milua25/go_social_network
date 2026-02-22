package main

import (
	"fmt"
	"math"
	"slices"
)

func maxProfit(prices []int) int {
	const (
		buyDay  = 2 // 1-based day index
		sellDay = 5 // 1-based day index
	)

	n := len(prices)

	// Validate days (1-based) and ordering.
	if buyDay < 1 || sellDay < 1 || buyDay > n || sellDay > n || sellDay < buyDay {
		return 0
	}

	return profitBetweenDays(buyDay, sellDay, prices)
}

func profitBetweenDays(buyDay, sellDay int, prices []int) int {
	buyPrice := prices[buyDay-1]   // convert to 0-based index
	sellPrice := prices[sellDay-1] // convert to 0-based index

	profit := sellPrice - buyPrice
	if profit < 0 {
		return 0
	}
	return profit
}

func main() {
	azz := []int{7, 1, 5, 3, 6, 4}
	println(maxProfit(azz))

	tests := []struct {
		name   string
		result bool
		nums   []int
	}{
		{name: "Example 1", nums: []int{1, 2, 3, 1}, result: true},
		{name: "Example 2", nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, result: true},
	}

	fmt.Println(tests)

}
func threeSum(nums []int) [][]int {
	length := len(nums)
	result := make([][]int, 0)
	if length < 3 {
		return result
	}

	slices.Sort(nums)

	for i := 0; i < length-2; i++ {
		first := nums[i]
		if i > 0 && first == nums[i-1] {
			continue
		}
		if first > 0 {
			break
		}

		target := -first
		left, right := i+1, length-1
		for left < right {
			twoSum := nums[left] + nums[right]
			switch {
			case twoSum > target:
				right--
			case twoSum < target:
				left++
			default:
				result = append(result, []int{first, nums[left], nums[right]})
				left++

				for left < right && nums[left] == nums[left-1] {
					left++
				}
			}
		}
	}

	return result
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if len(coins) == 0 {
		return -1
	}

	bestCount := math.MaxInt

	for _, coin := range coins {
		if coin <= 0 {
			continue
		}

		quotient := amount / coin
		remainder := amount % coin

		if remainder == 0 {
			if quotient < bestCount {
				bestCount = quotient
			}
			continue
		}

		if slices.Contains(coins, remainder) {
			candidate := quotient + 1
			if candidate < bestCount {
				bestCount = candidate
			}
		}
	}

	if bestCount == math.MaxInt {
		return -1
	}
	return bestCount
}
