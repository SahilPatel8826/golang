package main

import (
	"fmt"
)

// func Channel(c chan int) {
// 	for i := 0; i < 5; i++ {
// 		c <- i
// 	}

// }

// func main() {
// 	ch := make(chan int)

// 	go Channel(ch)

// 	for i := range ch {
// 		fmt.Println(i)
// 	}

// }

// func channel(c chan string, s []string) {

// 	for i := 0; i < len(s); i++ {
// 		c <- s[i]

// 	}
// 	close(c)

// }

// func main() {
// 	ch := make(chan string, 3)

// 	sl := []string{"a", "b", "c"}
// 	go channel(ch, sl)

// 	for msg := range ch {
// 		fmt.Println(msg)
// 	}
// }

/////////MY WRONG APPROACH
// func channel(a chan string, b chan string) {
// 	select {
// 	case v := <-a:
// 		fmt.Println("a", v)

// 	case v := <-b:
// 		fmt.Println("b", v)
// 	default:
// 		fmt.Println("fool")

// 	}

// }

// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	go channel(ch1, ch2)
// 	ch1 <- "Send A"
// 	ch2 <- "Send B"
// 	close(ch1)
// 	close(ch2)

// }

////////CORRECTED SOLUTION//////

// func channel(a, b chan string) {
// 	select {
// 	case v := <-a:
// 		fmt.Println("a:", v)
// 	case v := <-b:
// 		fmt.Println("b:", v)
// 	}
// }

// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		ch1 <- "Send A"
// 	}()

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch2 <- "Send B"
// 	}()

// 	channel(ch1, ch2)
// }
// 🧩 TASK 6: Multiple Goroutines Sending to One Channel

// Goal: Fan-in concept (basic)

// Task

// Start 3 goroutines

// Each sends its ID (1,2,3) to same channel

// Main goroutine reads 3 values

// Expected Learning

// Channels are thread-safe

// Order is not guaranteed
///////MY Attempt////////

// func main() {
// 	ch := make(chan int)

// 	for i := 0; i < 3; i++ {
// 		go func(ID int) {
// 			ch <- i
// 		}(i)
// 		close(ch)
// 	}

// 	for val := range ch {
// 		fmt.Println(val)
// 	}
// }

// SOLUTION

// func main() {
// 	ch := make(chan int)

// 	// start goroutines
// 	for i := 0; i < 3; i++ {
// 		go func(ID int) {
// 			ch <- ID
// 		}(i)
// 	}

//		// receive exactly 3 values
//		for i := 0; i < 3; i++ {
//			fmt.Println(<-ch)
//		}
//	}
// 🧩 TASK 8: Timeout Using time.After

// Goal: Avoid infinite blocking

// Task

// Receive from a channel

// Add timeout case using time.After(2 * time.Second)

// Print “timeout” if nothing arrives

// 📌 Interview concept:

// “Prevent goroutine leaks”

// func main() {
// 	ch := make(chan int)

// 	select {
// 	case val := <-ch:
// 		fmt.Println("Received:", val)

// 	case <-time.After(2 * time.Second):
// 		fmt.Println("Timeout: no data received")
// 	}
// }

func worker(c <-chan int) {
	for ran := range c {
		fmt.Printf("jobdone%d ", ran)
	}
}

func main() {
	ch := make(chan int)
	go worker(ch)

	for i := 0; i < 5; i++ {
		ch <- i
	}

	close(ch)

}
