package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	message := "Hello, World!"
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Println(encodedMessage)

}
