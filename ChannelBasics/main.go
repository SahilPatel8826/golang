package main

import (
	"fmt"
	"time"
)

//  urls := []string{
//     "https://www.google.com",
//     "https://www.github.com",
//     "https://www.stackoverflow.com",
//     "https://www.reddit.com",
//     "https://www.wikipedia.org",
//     "https://www.medium.com",
//     "https://www.bbc.com",
//     "https://www.cnn.com",
//     "https://golang.org",
//     "https://www.nytimes.com",
// }

// func dowloadPage(url string){
//     fmt.Printf("[%s] running")
// }

// func main(){
//    result:=make(chan string)
//    for _,url:=range in urls{
// 	  go downloadPage(url)

//	   }
//	}

//send
// func processMsg(msgChan chan string) {
// 	for msg := range msgChan { // keeps receiving until channel closed
// 		fmt.Println("Print", msg)
// 		time.Sleep(time.Second * 1)
// 	}
// }

//	func Receive(num chan int, a int, b int) {
//		numResult := a + b
//		num <- numResult
//	}
func sendEmail(emailChan chan string, done chan bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("Sending email to", email)
		time.Sleep(time.Second) // Simulate time taken to send email
	}
}

func main() {
	// num := make(chan int)
	// go Receive(num, 5, 10)
	// result := <-num
	// fmt.Println("Result:", result)
	// msgChan := make(chan string)
	// go processMsg(msgChan)
	// time.Sleep(time.Second * 2)
	// msgChan <- "Hello"
	// time.Sleep(time.Second * 2)
	// msgChan <- "World"
	// time.Sleep(time.Second * 2)
	// close(msgChan) // close the channel when done sending

	// for {

	// 	msgChan <- "Hello"
	// }

	emailChan := make(chan string, 100)
	done := make(chan bool)
	go sendEmail(emailChan, done)

	for i := 0; i < 5; i++ {

		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}
	fmt.Println("All emails sent")
	close(emailChan)
	<-done

}
