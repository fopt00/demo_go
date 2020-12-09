package goroutine

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func CrawlerEntry() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:\n\tgo run crawl.go <url>")
		os.Exit(1)
	}

	workList := make(chan []string)
	seen := make(map[string]bool)

	var n int
	n++

	go func() {
		workList <- os.Args[1:]
	}()

	for ; n > 0; n-- {
		list := <-workList
		for _, url := range list {
			if seen[url] {
				continue
			}
			n++
			seen[url] = true
			go func(url string) {
				workList <- crawl(url)
			}(url)
		}
	}
}

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	urls, err := ExtractLinks(url)
	<-tokens
	if err != nil {
		log.Print(fmt.Sprintf("\x1b[31m%v\x1b[0m", err))
	}
	return urls
}

func ExtractLinks(seed string) ([]string, error) {
	var urls []string

	pattern := regexp.MustCompile(`<a\b[^>]+\bhref="([^"]*)"[^>]*>[\s\S]*?</a>`)
	resp, err := http.Get(seed)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Println("exception happened when response close")
		}
	}()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := pattern.FindAllStringSubmatch(string(buf), -1)

	for _, e := range res {
		if len(e) != 2 {
			continue
		}

		if strings.HasPrefix(e[1], "http") {
			urls = append(urls, e[1])
		}
	}
	return urls, nil
}
