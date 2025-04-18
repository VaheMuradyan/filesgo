package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := "example.txt"

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Full file content:")
	fmt.Println(string(content))
}
