package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// Compile and Execute: go build -o goexecutable && ./goexecutable
// Get help and parameter details: go build . && ./goexecutable -h

type problem struct {
	question string
	answer string
}

func main()  {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for quiz in seconds")
	flag.Parse()

	problems := processFile(*csvFilename)

	startTest(problems, *timeLimit)
}

func processFile(csvFilename string) []problem {
	file, err := os.Open(csvFilename)
	if err != nil {
		exit(err, fmt.Sprintf("Failed to open the CSV File: %s", csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(err, "Failed to parse provided CSV file")
	}

	return parseLines(lines)
}

func parseLines(lines [][]string) []problem {
	ret := make ([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

func startTest(problems []problem, timeLimit int) {
	var answer string
	correct := 0
	answerCh := make(chan string)
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)

		go func() {
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.answer {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func exit(err error, msg string) {
	fmt.Println(err)
	fmt.Println(msg)
	os.Exit(1)
}