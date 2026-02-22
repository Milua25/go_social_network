package main

import (
	"fmt"
	"time"
)

func main() {
	// current local time
	fmt.Println(time.Now())

	specificTime := time.Date(2024, time.June, 30, 0, 0, 0, 0, time.UTC)
	fmt.Println(specificTime)

	// ParseTime
	parsedTime, err := time.Parse("2006-01-02", "2020-05-01")
	if err != nil {
		panic(err)
	} // Mon Jan 2 2006 15:04:05 MST
	fmt.Println("Parsed Time:", parsedTime)

	nowTime := time.Now()
	unixTime := nowTime.Unix()

	fmt.Println("Unix time:", unixTime)

	layout := "2006-01-02T15:04:05Z07:00"
	str := "2025-07-04T14:30:18Z"

	t, err := time.Parse(layout, str)
	if err != nil {
		panic(err)
	}
	fmt.Println("Parsed Time:", t)
}
