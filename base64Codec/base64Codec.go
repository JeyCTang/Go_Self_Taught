package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// the string need to be proceeded
	message := "Away from keyboard. https://golang.org/"

	// encoded message
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(message))
	// print the encoding process is finished
	fmt.Println(encodedMessage)
	// decode the message
	data, err := base64.StdEncoding.DecodeString(encodedMessage)
	// process the error
	if err != nil{
		fmt.Println(err)
	} else{
		// print decoded data
		fmt.Println(string(data))
	}
}
