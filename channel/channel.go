package main

import (
	"fmt"
	"time"
)

func main() {
	channelWithoutCache()
	channelWithCache()
}

func channelWithoutCache() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Second)

		ch <- "hello,msg for channel"
	}()

	msg := <-ch

	fmt.Print(msg)

}

func channelWithCache() {
	ch := make(chan string, 1)

	go func() {

		ch <- "hello, first msg from channel"

		time.Sleep(time.Second)

		ch <- "hello, second msg from channel"

	}()

	time.Sleep(2 * time.Second)

	msg := <- ch

	fmt.Println(time.Now().String() + msg)

	msg = <- ch

	fmt.Println(time.Now().String() + msg)
}
