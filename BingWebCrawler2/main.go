package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var bingDomains = map[string]string{
	"us": "&cc=US",
	"ca": "&cc=CA",
	"mx": "&cc=MX",
	"br": "&cc=BR",
	"ar": "&cc=AR",
	"cl": "&cc=CL",
	"co": "&cc=CO",
	"pe": "&cc=PE",
	"ve": "&cc=VE",

	"uk": "&cc=GB",
	"ie": "&cc=IE",
	"fr": "&cc=FR",
	"de": "&cc=DE",
	"es": "&cc=ES",
	"pt": "&cc=PT",
	"it": "&cc=IT",
	"nl": "&cc=NL",
	"be": "&cc=BE",
	"ch": "&cc=CH",
	"at": "&cc=AT",
	"se": "&cc=SE",
	"no": "&cc=NO",
	"dk": "&cc=DK",
	"fi": "&cc=FI",
	"pl": "&cc=PL",
	"cz": "&cc=CZ",
	"sk": "&cc=SK",
	"hu": "&cc=HU",
	"ro": "&cc=RO",
	"bg": "&cc=BG",
	"hr": "&cc=HR",
	"si": "&cc=SI",
	"rs": "&cc=RS",
	"ua": "&cc=UA",
	"ru": "&cc=RU",

	"in": "&cc=IN",
	"pk": "&cc=PK",
	"bd": "&cc=BD",
	"lk": "&cc=LK",
	"np": "&cc=NP",

	"cn": "&cc=CN",
	"hk": "&cc=HK",
	"tw": "&cc=TW",
	"jp": "&cc=JP",
	"kr": "&cc=KR",
	"sg": "&cc=SG",
	"my": "&cc=MY",
	"id": "&cc=ID",
	"th": "&cc=TH",
	"ph": "&cc=PH",
	"vn": "&cc=VN",

	"au": "&cc=AU",
	"nz": "&cc=NZ",

	"za": "&cc=ZA",
	"ng": "&cc=NG",
	"ke": "&cc=KE",
	"eg": "&cc=EG",
	"ma": "&cc=MA",

	"sa": "&cc=SA",
	"ae": "&cc=AE",
	"qa": "&cc=QA",
	"kw": "&cc=KW",
	"om": "&cc=OM",
	"il": "&cc=IL",
	"tr": "&cc=TR",
}

type SearchResult struct {
	ResultRank  int
	ResultTitle string
	ResultURL   string
	ResultDesc  string
}

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
}

func randomUserAgent() string {
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

// func main
// bing scrape
// build URL
// scrape client request
// parse result

// print to screen
func buildURL(searchTerm, country string, pages, count int) ([]string, error) {
	// if topic == "" {
	// 	topic = "golang"
	// }
	// if count <= 0 {
	// 	count = 1
	// }
	// baseURL := "https://www.bing.com/search?"

	// params := url.Values{}
	// params.Add("q", topic)
	// params.Add("count", strconv.Itoa(count))

	// return baseURL + params.Encode(), nil
	toScrap := []string{}
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	if countryCode, found := bingDomains[country]; found {
		for i := 0; i < pages; i++ {
			first := firstParameter(i, count)
			scrapeURL := fmt.Sprintf("https://bing.com/search?q=%s&count=%d&first=%d%s", searchTerm, count, first, countryCode)
			toScrap = append(toScrap, scrapeURL)
		}
	} else {
		return nil, fmt.Errorf("Invalid country code")
	}
	return toScrap, nil

}

func firstParameter(number, count int) int {
	if number == 0 {
		return number + 1
	}
	return number*count + 1
}
func getScrapeClient(proxyString interface{}) *http.Client {
	switch v := proxyString.(type) {
	case string:
		proxyUrl, _ := url.Parse(v)
		return &http.Client{
			Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}

	default:
		return &http.Client{}
	}
}

func scrapeClientRequest(searchURL string, proxyString interface{}) (*http.Response, error) {
	// client := http.Client{
	// 	Timeout: 10 * time.Second,
	// }
	// req, err := http.NewRequest("GET", url, nil)
	// req.Header.Set("User-Agent", randomUserAgent())
	// if err != nil {
	// 	return nil, err
	// }
	// res, err := client.Do(req)
	// if err != nil {
	// 	return nil, err
	// }
	// return res, nil\\
	client := getScrapeClient(nil)
	req, err := http.NewRequest("GET", searchURL, nil)
	req.Header.Set("User-Agent", randomUserAgent())

	res, err := client.Do(req)
	if res.StatusCode != 200 {
		err := fmt.Errorf("Non 200 response code received")
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return res, nil

}

func bingResultParser(response *http.Response, rank int) ([]SearchResult, error) {
	results := []SearchResult{}
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return nil, err
	}
	sel := doc.Find("li.b_algo")
	rank++
	for i := range sel.Nodes {
		item := sel.Eq(i)
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		titleTag := item.Find("h2")
		descTag := item.Find("p")
		desc := descTag.Text()
		title := titleTag.Text()
		link = strings.Trim(link, " ")
		if link != "" && link != "#" && !strings.HasPrefix(link, "/") {
			result := SearchResult{
				rank,
				link,
				title,
				desc,
			}
			results = append(results, result)
			rank++
		}
	}
	return results, nil
}
func bingScrape(searchTerm, country string, proxyString interface{}, pages, backoff, count int) ([]SearchResult, error) {
	results := []SearchResult{}

	bingPages, err := buildURL(searchTerm, country, pages, count)
	if err != nil {
		return nil, err
	}

	for _, page := range bingPages {
		rank := len(results)
		res, err := scrapeClientRequest(page, proxyString)
		if err != nil {
			return nil, err
		}
		data, err := bingResultParser(res, rank)
		if err != nil {
			return nil, err
		}
		for _, result := range data {
			results = append(results, result)

		}
		time.Sleep(time.Duration(backoff) * time.Second)
	}
	return results, nil
}

func main() {
	// reader := bufio.NewReader(os.Stdin)

	// fmt.Println("Enter the topic to serach on Bing: ")
	// topic, _ := reader.ReadString('\n')
	// topic = strings.TrimSpace(topic)

	// fmt.Print("Enter count: ")
	// countStr, _ := reader.ReadString('\n')
	// countStr = strings.TrimSpace(countStr)

	// count, _ := strconv.Atoi(countStr)

	// bingURL := buildURL(topic, count)

	// fmt.Println("🔍 Bing Search URL:")
	// fmt.Println(bingURL)

	// p := DefaultParser{}
	// results := ScrapeSitemap(bingURL, p, 10)
	// for _, res := range results {
	// 	fmt.Println(res)
	// }

	res, err := bingScrape("tomato", "us", nil, 2, 30, 30)
	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}

	} else {
		fmt.Println(err)

	}
}
