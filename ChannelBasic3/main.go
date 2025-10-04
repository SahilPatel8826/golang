package main

import (
	"fmt"
	"time"
)

// Worker function — reads from jobChan
func ProcessURL(id int, jobChan <-chan string) {
	for url := range jobChan {
		fmt.Printf("Worker %d processing: %s\n", id, url)
		time.Sleep(time.Second) // simulate network call
		fmt.Printf("Worker %d finished: %s\n", id, url)
	}
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://www.reddit.com",
		"https://www.wikipedia.org",
		"https://www.medium.com",
		"https://www.bbc.com",
		"https://www.cnn.com",
		"https://golang.org",
		"https://www.nytimes.com",
	}

	jobChan := make(chan string, len(urls))

	// Start 5 workers
	for i := 1; i <= 5; i++ {
		go ProcessURL(i, jobChan)
	}

	// Send jobs to channel
	for _, url := range urls {
		jobChan <- url
	}

	close(jobChan) // important: close when done sending

	time.Sleep(5 * time.Second) // wait for workers to finish (simple wait)
}
