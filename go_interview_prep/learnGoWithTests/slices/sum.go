package main

import "fmt"

func Sum(numbers ...int) int {
	var sum int
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

func SumAll(sl1, sl2 []int) (result []int) {
	result = []int{Sum(sl1...), Sum(sl2...)}
	return
}

func SumTails(sl1, sl2 []int) []int {
	var result []int
	if len(sl1) == 0 || len(sl2) == 0 {
		return result
	}
	sl1 = sl1[1:]
	sl2 = sl2[1:]
	return SumAll(sl1, sl2)
}

func UpdateSlice(s1, s2 []int) []int {
	return append(s1, s2...)
}

func CreateSlice() {

	twoD := make([][]int, 3) // [][]int{[0],[0], [0] }

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)
}

func main() {
	azz := []int{1, 2, 3}
	zzz := []int{4, 5, 6}
	fmt.Println(SumAll(azz, zzz))

	CreateSlice()

	process()
}

func process() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered in f", err)
		}
	}()
	fmt.Println("Start processing")
	panic("Something went wrong")
	fmt.Println("End processing")
}

// panic

// defer - all ways run even when panic occurs

// recover
