package main

import "fmt"

func main() {
	messages := make(chan string, 2) // 2 banyanykan channel yang dibuat dalam string
	messages <- "buffered"
	messages <- "channel"
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

// $ go run channel-buffering.go
// buffered
// channel
