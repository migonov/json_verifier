package main

import (
	"fmt"
	"os"

	"github.com/migonov/json_verifier/verifier"
)

const FileName = "data.json"

func main() {
	data, err := os.ReadFile(FileName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("read %s file content \n%s\n", FileName, string(data))

	isValid := verifier.Verify([]byte(data))
	if isValid {
		fmt.Printf("%s is valid\n", FileName)
	} else {
		fmt.Printf("%s is invalid\n", FileName)
	}
}
