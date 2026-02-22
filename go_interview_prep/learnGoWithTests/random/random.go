package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randNew := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Println("Random Value:", randNew.Intn(100)+1)

}
