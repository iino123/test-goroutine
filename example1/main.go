package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for {
	// 	// 第一引数に <-c があることで、channelからのメッセージを受け取るまで処理がブロックされている。
	// 	// TODO: このgoキーワードを無くした時の挙動について
	// 	go checkLink(<-c, c)
	// }

	// 上記に同じ意味合いになる
	// range <channel>とする書き方
	for l := range c {
		// time.Sleep(5 * time.Second) NOTE: ここでブロックしてしまうとメインのgo-routineをブロックしてしまうので、要件を満たせない
		// go checkLink(l, c)

		// TODO: これでは不具合になる、変数lがなんかおかしくなる
		// fanction literal (jsで言うところの、無名関数の即時実行)
		// go func() {
		// 	time.Sleep(5 * time.Second)
		// 	checkLink(l, c)
		// }()

		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)

	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
