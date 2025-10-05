package main

import (
	"fmt"
	"time"
)

func MsgGenerator(msgChan chan<- string, count *int) {
	*count++
	for {
		msg := fmt.Sprintf("Message %d", *count)
		msgChan <- msg
		time.Sleep(300 * time.Millisecond) // simulate delay
	}
}

func worker(id int, msgChan <-chan string) {
	for msg := range msgChan {
		fmt.Printf("Worker %d received: %s\n", id, msg)
	}
}

func main() {
	msgChan := make(chan string, 10)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	counter := 0
	// Start 5 workers (each waits for one message)
	for i := 1; i <= 3; i++ {
		go worker(i, msgChan)
	}

	go func() {
		for range ticker.C {

			go MsgGenerator(msgChan, &counter)
		}
	}()

	time.Sleep(10 * time.Second) // wait before program exits
	close(msgChan)
	fmt.Println("Ticker stopped.Program exiting...")
}
