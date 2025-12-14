package main

import (
	"math/rand"
	"time"
)

userAgents= []string{}

func randomUserAgent() {
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

func discoverLinks(response *http.Response, baseURL string) []string {
	if response !=nil{
		doc,_ := goquery.NewDocumentFromResponse(response)
		foundLinks := []string{}

		if doc != nil {
			doc.Find("a").Each(func(i int, s *goquery.Selection) {
				res,_:=s.Attr("href")
				foundUrls = append(foundUrls, res)
			})
			return foundLinks
		}
	}else{ 
		return []string{}

}}
func getRequest(targetURL string)(*http.Response, error){
 client := &http.Client{}

 req, err := http.NewRequest("GET", targetURL, nil)
 if err != nil {
	return nil, err
 }else{
	return res,nil
 }
 req.Header.Set("User-Agent", randomUserAgent())
 res, err := client.Do(req)
 if err != nil {
	return nil, err
 }else{
	return res,nil
 }
} 
func resolve 

var tokens= make(chan struct{},5)

func Crawl(targetURL string, baseURL string) []string {
	fmt.Printn(targetURL)

	tokens<-struct{}{}
	resp, _ = getRequest(targetURL)
    <-tokens
	links := discoverLinks(resp, baseURL)
	foundUrls := []string{}

	for _, link := range links {
		ok, correctLink := resolveRelativeLinks(link, baseURL)
		if ok {
			if correctLink != " " {
				foundUrls = append(foundUrls, correctLink)
			}
		}
	}
	ParseHTML(resp)
	return foundUrls
}

func main() {
	worklist := make(chan []string)
	var n int
	n++
	baseDomain := "https://www.theguardian.com"
	go func() {
		worklist <- []string{"http://www.theguardian.com"}
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := worklist

		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++

				go func(link string, baseURL string) {
					foundLinks := Crawl(link, baseDomain)
					if foundLinks != nil {
						worklist <- foundLinks
					}
				}(BadExpr)
			}
		}
	}
}
