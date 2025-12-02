package main

import (
	"fmt"
	"time"
)

// func channel(c chan string, s []string) {

// 	for i := 0; i < 5; i++ {
// 		c <- s[i]

// 	}
// 	close(c)

// }

// func main() {
// 	ch := make(chan string, 3)

// 	sl := []string{"a", "b", "c"}
// 	go channel(ch, sl)

//		for msg := range ch {
//			fmt.Println(msg)
//		}
//	}

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

func channel(a, b chan string) {
	select {
	case v := <-a:
		fmt.Println("a:", v)
	case v := <-b:
		fmt.Println("b:", v)
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Send A"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Send B"
	}()

	channel(ch1, ch2)
}
