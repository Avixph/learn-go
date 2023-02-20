package main

import (
	"time"
)

func main() {
	urls := []string{
		"https://www.google.com/", "https://www.facebook.com/", "https://stackoverflow.com/", "https://go.dev/", "https://www.amazon.com/",
	}

	c := make(chan string)

	for _, url := range urls {
		// checkUrl(url)
		go checkUrl(url, c)
	}
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// for i := 0; i < len(urls); i++ {
	// 	// fmt.Println(<-c)
	// 	go checkUrl(<-c)
	// }
	// for {
	// 	// fmt.Println(<-c)
	// 	go checkUrl(<-c, c)
	// }
	for u := range c {
		// time.Sleep(5 * time.Second)
		// go checkUrl(l, c)
		go func(url string) {
			time.Sleep(5 * time.Second)
			checkUrl(url, c)
		}(u)
	}
}
