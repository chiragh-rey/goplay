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
// Get help and parameter details: go build . && ./quizgame -h

func main()  {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(err, fmt.Sprintf("Failed to open the CSV File: %s", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(err, "Failed to parse provided CSV file")
	}

	problems := parseLines(lines)
	var answer string
	correct := 0
	answerCh := make(chan string)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)

		go func() {
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
			case <-timer.C:
				fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
				return
			case answer := <-answerCh:
				if answer == p.a {
					correct++
				}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make ([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

type problem struct {
	q string
	a string
}

func exit(err error, msg string) {
	fmt.Println(err)
	fmt.Println(msg)
	os.Exit(1)
}