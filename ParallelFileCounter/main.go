package main

import (
	"fmt"
	"os"
	"strings"
)

func ReadAllFiles(countChan chan int, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("file is not opening", err)
		return
	}
	defer file.Close()

	data := make([]byte, 1024)
	count, err := file.Read(data)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	content := string(data[:count])

	words := strings.Fields(content)

	wordCount := len(words)

	fmt.Println("File content %s:", filename)
	fmt.Println("\nTotal words:", wordCount)

	countChan <- wordCount
}

func main() {

	filenames := []string{
		"about.txt",
		"golang.txt",
		"history.txt"}

	countChan := make(chan int, 3)
	for _, names := range filenames {
		go ReadAllFiles(countChan, names)
	}

	totalWords := 0

	// Collect results from each goroutine
	for range filenames {
		count := <-countChan
		totalWords += count
	}

	fmt.Println("Total words in all files:", totalWords)

}
