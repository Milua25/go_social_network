package iteration

import (
	"fmt"
	"os"
	"strings"
)

const repeatCount = 5

func Repeat(value string, count int) string {
	//return fmt.Sprintf("%s%s%s%s%s", value, value, value, value, value)
	var repeated strings.Builder
	if count == 0 {
		count = repeatCount
	}

	for i := 0; i < count; i++ {
		_, err := repeated.WriteString(value)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
	}
	return repeated.String()
}
