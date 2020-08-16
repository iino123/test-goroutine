package main

import "fmt"

// Printlnされる順番に注意
// receiving from channelが一番最初に表示されることもある。
// しかし、greeting := <-cの部分でチャネルから値が返ってくるのを待つ
func main() {
	c := make(chan string)

	go func(input chan string) {
		fmt.Println("sending to channel")
		input <- "hello"
	}(c)

	fmt.Println("receiving from channel")
	greeting := <-c
	fmt.Println("greeting recieved")
	fmt.Println(greeting)
}
