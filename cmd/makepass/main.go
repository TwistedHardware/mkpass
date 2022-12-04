package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/twistedhardware/mkpass"
)

func main() {
	length := 5
	passType := "words"

	if len(os.Args) > 1 {
		// Check if the user is looking for help
		if strings.Contains(os.Args[1], "help") {
			help := `Usage: mkpass length [OPTIONS]
  --base64    Generates a base64 password

examples:
 mkpass 6
 This will return a password of length 6 words

 mkpass 24 --base64
 This will return a password of length 24 base64 charters
`
			fmt.Println(help)
			return
		}

		// check for password type
		if os.Args[1] == "--base64" {
			passType = "base64"
			length = 15
		} else if len(os.Args) > 2 && os.Args[2] == "--base64" {
			passType = "base64"
			length = 15
		}

		// check if there is a length
		lengthRaw := os.Args[1]
		if val, err := strconv.Atoi(lengthRaw); err == nil && val > 0 {
			length = val
		}
	}

	if passType == "base64" {
		fmt.Println(mkpass.GenerateBase64(length))
		return
	}

	fmt.Println(mkpass.GenerateDicewarePassword(length))
}
