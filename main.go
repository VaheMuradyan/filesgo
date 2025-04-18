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

	ReadRuny(filePath)
}

func ReadRuny(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	charactersToRead := 100

	for charactersToRead > 0 {
		char, _, err := reader.ReadRune()
		if err != nil {
			break
		}

		fmt.Println(string(char))
		charactersToRead--
	}

	fmt.Println("...")
}

func ReadLineAndRuny(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Seek(0, 0) // senc pointer@ hetem berum amenaskizb
}
