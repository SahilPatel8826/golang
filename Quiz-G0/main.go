package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func ProblemParser(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func ParseProblem(lines [][]string) []problem {
	r := make([]problem, len(lines))

	for i := 0; i < len(lines); i++ {
		r[i] = problem{Question: lines[i][0], Answer: lines[i][1]}
	}

	return r
}

func main() {
	// Step 1: Parse CSV file
	lines, err := ProblemParser("quiz.csv")
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// Step 2: Convert CSV rows into []problem
	problems := ParseProblem(lines)

	// Step 3: Start quiz
	score := 0
	reader := bufio.NewReader(os.Stdin)

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.Question)

		// Read user input
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Check answer
		if input == p.Answer {
			fmt.Println("✔ Correct!")
			score++
		} else {
			fmt.Println("✘ Wrong! Correct answer =", p.Answer)
		}

		fmt.Println()
	}

	// Step 4: Show final score
	fmt.Printf("You scored %d out of %d\n", score, len(problems))
}

type problem struct {
	Question string
	Answer   string
}
