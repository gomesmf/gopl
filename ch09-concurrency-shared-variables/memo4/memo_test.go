package memo

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func incomingURLs() []string {
	return []string{
		"https://golang.org",
		"https://google.com",
		"https://play.golang.org",
		"https://gopl.io",
		"https://golang.org",
		"https://google.com",
		"https://play.golang.org",
		"https://gopl.io",
	}
}

func TestHttpGetBodySequential(t *testing.T) {
	m := New(httpGetBody)
	for _, url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}

}

func TestHttpGetBodyConcurrent(t *testing.T) {
	m := New(httpGetBody)
	var n sync.WaitGroup
	for _, url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
}
