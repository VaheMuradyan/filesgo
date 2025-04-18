package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	filePath := "example.txt"
	totalSum := 0
	totalCount := 0

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		val, _ := strconv.Atoi(line)
		totalCount++
		totalSum += val
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if totalCount > 0 {
		fmt.Println(totalSum / totalCount)
	}
}
