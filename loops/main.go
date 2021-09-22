package main

import "fmt"
import "time"

func main() {


	x:=10
	for x<100 {
		fmt.Println(x)
		x=x+10
	}

	loopTimer:=time.NewTimer(time.Second*5)
	for{
		fmt.Println("INFINITE")

		<-loopTimer.C
		break

	}
}
