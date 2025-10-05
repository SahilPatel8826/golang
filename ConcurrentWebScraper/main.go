package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
)

// ExtractHeadlines fetches and sends all h1, h2, h3 headlines from a URL
func ExtractHeadlines(headChan chan<- map[string][]string, url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching:", url, err)
		headChan <- map[string][]string{url: {"[Error fetching page]"}}
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		fmt.Println("Error parsing HTML from:", url, err)
		headChan <- map[string][]string{url: {"[Error parsing HTML]"}}
		return
	}

	var headlines []string
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && (n.Data == "h1" || n.Data == "h2" || n.Data == "h3") {
			if n.FirstChild != nil {
				text := strings.TrimSpace(n.FirstChild.Data)
				if text != "" {
					headlines = append(headlines, text)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(doc)

	headChan <- map[string][]string{url: headlines}
}

func main() {
	urls := []string{
		"https://www.bbc.com",
		"https://www.cnn.com",
		"https://www.nytimes.com",
	}

	headChan := make(chan map[string][]string)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			for _, url := range urls {
				go ExtractHeadlines(headChan, url)
			}
		}
	}()

	// Collect and print results continuously
	go func() {
		for data := range headChan {
			for url, headlines := range data {
				fmt.Printf("\n============================\n")
				fmt.Printf("🔹 Headlines from %s:\n", url)
				fmt.Printf("============================\n")

				if len(headlines) == 0 {
					fmt.Println("No headlines found.")
				} else {
					for i, h := range headlines {
						fmt.Printf("%d. %s\n", i+1, h)
					}
				}
			}
		}
	}()

	time.Sleep(20 * time.Second)
	fmt.Println("\nTicker stopped. Program exiting...")

}
