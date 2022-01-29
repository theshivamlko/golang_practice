package main

import (
	"fmt"
	//"time"
)

// Buffer Channel is Queue

var done chan bool = make(chan bool)

func main() {
	fmt.Println("Start")
	go print("GO")
	<-done
	print("NORMAL")

}

func print(str string) {
	for i := 0; i < 200; i++ {
		fmt.Println(i, str)
	}

	if str == "GO" {
		fmt.Println("GO========")
		done <- true
	}
	//	time.Sleep(time.Millisecond * 1)
}
