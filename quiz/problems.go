package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type line struct {
	question, answer string
}

func readFile(fileName string) []line {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()

	if err != nil {
		log.Fatal(err)
	}
	var newArr []line
	for _, parsedLine := range lines {
		newArr = append(newArr, line{
			question: parsedLine[0],
			answer:   parsedLine[1],
		})
	}
	return newArr
}

func parseLine(qa line) bool {
	fmt.Print(qa.question)
	fmt.Print(" = ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	userAnswer := input.Text()
	return userAnswer == qa.answer
}

func main() {
	data := readFile("problems.csv")

	var counter int
	fmt.Printf("Try to answer following %v questions!\n", len(data))
	for _, line := range data {
		success := parseLine(line)
		if success {
			counter++
		}
	}

	fmt.Println("You got", counter, "answer(s) correctly!")
}
