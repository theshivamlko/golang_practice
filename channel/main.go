package main

import "fmt"

func writeMsg(message chan string) {
	message <- "Hello"
}

func main() {

	fmt.Println("Channel")
	message := make(chan string)
	go writeMsg(message)

	fmt.Println("Greeting ", <-message)
	close(message)

}
