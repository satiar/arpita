package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go func(channel chan string) {
		test2(channel)
	}(channel)

	time.Sleep(10 * time.Millisecond)
	close(channel)
}

func test2(channel chan string) {
	for {
		select {
		case _, ok := <-channel:
			if !ok {
				fmt.Println(" I GOT THE EXIT MESSAGE")
				return
			}
		default:
			fmt.Println("I AM IN DEFAULT")
		}
	}
}
