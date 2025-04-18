package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	WriteFile("output.txt")
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

func WriteFile(filePath string) {
	//grumenq elaci vren jnjvuma hin@
	err := os.WriteFile(filePath, []byte(fmt.Sprintf("%d", 22)), 0664)
	if err != nil {
		log.Fatal(err)
	}

	//grumenq hin fili takic
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString("Taki toxicenq avelacnum\n")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fmt.Println("exav")

}

func ReadTable(filepath string) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	var data [][]int

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.Fields(line)
		row := make([]int, len(parts))
		for i, part := range parts {
			row[i], _ = strconv.Atoi(part)
		}

		data = append(data, row)
	}

	for _, row := range data {
		fmt.Println(row)
	}
}
