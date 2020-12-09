package goroutine

import (
	"fmt"
	"log"
	"os"
)

func SpiderEntry() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:\n\t parameters should contain some urls")
		os.Exit(1)
	}

	workList := make(chan []string)
	unseenLinks := make(chan string)
	seen := make(map[string]bool)

	var n int
	n++
	go func() {
		workList <- os.Args[1:]
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for url := range unseenLinks {
				urls := spider(url)
				go func() { workList <- urls }()
			}
		}()
	}

	for list := range workList {
		for _, url := range list {
			if seen[url] {
				continue
			}
			unseenLinks <- url
		}
	}

}

func spider(url string) []string {
	fmt.Println(url)
	urls, err := ExtractLinks(url)
	if err != nil {
		log.Print(fmt.Sprintf("\x1b[31m%v\x1b[0m", err))
	}
	return urls
}
