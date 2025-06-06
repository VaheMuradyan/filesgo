package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Person struct {
	Name       string
	Age        int
	Occupation string
}

func main() {
	// filePath := "example.txt"
	// totalSum := 0
	// totalCount := 0

	// file, err := os.Open(filePath)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// scanner := bufio.NewScanner(file)

	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	val, _ := strconv.Atoi(line)
	// 	totalCount++
	// 	totalSum += val
	// }

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// if totalCount > 0 {
	// 	fmt.Println(totalSum / totalCount)
	// }

	// ReadRuny(filePath)

	// WriteFile("output.txt")

	// ReadFromFileInStruct("data.txt")

	// WriteCSV("csvout.csv")

	//          ReadJson("data.json")
	ReadJson2("data2.json")

	// r := gin.Default()

	// r.GET("/process", func(c *gin.Context) {
	// 	var data map[string]interface{}
	// 	if err := c.BindJSON(&data); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	// Extract and print the school name
	// 	schoolName := data["school"].(string)
	// 	fmt.Println("School Name:", schoolName)

	// 	// Extract and print the city
	// 	location := data["location"].(map[string]interface{})
	// 	city := location["city"].(string)
	// 	fmt.Println("School's City:", city)

	// 	// Extract and print the name of the second student
	// 	students := data["students"].([]interface{})
	// 	secondStudent := students[1].(map[string]interface{})
	// 	secondStudentName := secondStudent["name"].(string)
	// 	fmt.Println("Name of the Second Student:", secondStudentName)

	// 	// Return response
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"school":         schoolName,
	// 		"city":           city,
	// 		"second_student": secondStudentName,
	// 	})
	// })

	// r.Run(":8080")

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

func ReadFromFileInStruct(filepath string) {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	var people []Person

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		age, _ := strconv.Atoi(parts[1])
		people = append(people, Person{
			Name:       parts[0],
			Age:        age,
			Occupation: parts[2],
		})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Parsed People Data:")
	for _, person := range people {
		fmt.Printf("Name: %s, Age: %d, Occupation: %s\n", person.Name, person.Age, person.Occupation)
	}

}

func ReadCSV(filepath string) {
	var names []string // Slice to store names

	// Open the CSV file
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records, skipping the header
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Iterate over the records
	for _, record := range records[1:] { // Skip header on index 0
		// TODO: Extract the name from 'record'
		name := record[0]

		// TODO: Append the extracted name to the 'names' slice
		names = append(names, name)
	}

	// TODO: Sort the 'names' slice in alphabetical order
	// - You can use the sort.Strings() method for that
	sort.Strings(names)

	// TODO: Print each name from the sorted 'names' slice
	fmt.Println(names)
}

func WriteCSV(filepath string) {
	data := []Person{
		{Name: "John", Age: 28, Occupation: "Engineer"},
		{Name: "Alice", Age: 34, Occupation: "Doctor"},
		{Name: "Bob", Age: 23, Occupation: "Artist"},
	}

	csvFile, err := os.Create(filepath)

	if err != nil {
		log.Fatal(err)
	}

	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)

	defer writer.Flush()

	writer.Write([]string{"Name", "Age", "Occupation"})

	for _, person := range data {
		writer.Write([]string{person.Name, fmt.Sprint(person.Age), person.Occupation})
	}

	fmt.Println("Data successfully written to output.csv (comma-separated).")
}

func WriteFileTable(filepath string) {
	data := []Person{
		{Name: "John", Age: 28, Occupation: "Engineer"},
		{Name: "Alice", Age: 34, Occupation: "Doctor"},
		{Name: "Bob", Age: 23, Occupation: "Artist"},
	}

	txtLines := []string{"Name Age Occupation"} // Header with spaces
	for _, person := range data {
		// Add each person's data as a space-separated line
		txtLines = append(txtLines, person.Name+" "+fmt.Sprint(person.Age)+" "+person.Occupation)
	}

	// Write the list of lines to the TXT file
	err := os.WriteFile(filepath, []byte(strings.Join(txtLines, "\n")), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Printf("Data successfully written to %s (space-separated).\n", filepath)
	}

	// Read and print the content of the file as a single string
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading file:", err)
	} else {
		fmt.Println("File content:")
		fmt.Println(string(fileContent))
	}

}

func ReadJson(filePath string) {
	// Read the entire content of the JSON file into a byte slice
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the JSON content into a map
	var data map[string]interface{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Parsed JSON Data:")
	// prettyJSON, err := json.MarshalIndent(data, "", " ")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Extract and print the school name
	schoolName := data["school"].(string)
	fmt.Println("School Name:", schoolName)

	// Extract and print the city
	location := data["location"].(map[string]interface{})
	city := location["city"].(string)
	fmt.Println("School's City:", city)

	// Extract and print the name of the second student
	students := data["students"].([]interface{})
	secondStudent := students[1].(map[string]interface{})
	secondStudentName := secondStudent["name"].(string)
	fmt.Println("Name of the Second Student:", secondStudentName)
}

func ReadJson2(filePath string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var data map[string]interface{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Retrieve the "departments" array from the JSON data and ensure it is cast to []interface{}
	departments, ok := data["departments"].([]interface{})
	if !ok {
		fmt.Println("Error")
	}

	for _, department := range departments {
		dept := department.(map[string]interface{})
		employees, ok := dept["employees"].([]interface{})
		if !ok {
			fmt.Println("Error")
		}
		for _, employee := range employees {
			emp := employee.(map[string]interface{})
			fmt.Println(emp["name"])
		}

	}

	// TODO: Check if the departments slice is not nil

}

func readJSON(path string) (map[string]interface{}, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data map[string]interface{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, fmt.Errorf("error")
	}

	return data, nil
}
