package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
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

func createTimer(timeLimit int) *time.Timer {
	return time.NewTimer(time.Duration(timeLimit) * time.Second)
}

func handleTimer(timer time.Timer, counter *int) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("\n\n\n*** You ran out of time!***\n\n")
				printAnswersCount(*counter)
			}
		}()

		<-timer.C
		panic("You ran out of time!")
	}()
}

func printAnswersCount(counter int) {
	log.Fatal("You got ", counter, " answer(s) correctly!")
}

func checkUserReady(data []line, timeLimit int) {
	fmt.Printf("Try to answer following %v questions!\n", len(data))
	fmt.Println("Press Enter key to start the test: you've got", timeLimit, "seconds to answer all the questions!")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}

func parseFlags() (*string, *int) {
	csvFileName := flag.String("csv", "problems.csv", ".csv filename for questions and answers, comma separated")
	timeLimit := flag.Int("limit", 30, "the time limit for quic completion in seconds")
	flag.Parse()
	return csvFileName, timeLimit
}

func main() {
	csvFileName, timeLimit := parseFlags()

	data := readFile(*csvFileName)

	initCount := 0
	counter := &initCount

	checkUserReady(data, *timeLimit)

	timer := createTimer(*timeLimit)
	handleTimer(*timer, counter)

	for _, line := range data {
		success := parseLine(line)
		if success {
			*counter++
		}
	}

	printAnswersCount(*counter)
}
