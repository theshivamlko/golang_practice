package main

import (
	"fmt"
	"sync"
)


var greet string
var howdyDone chan bool
var mutex=&sync.Mutex{}

func main() {
	howdyDone=make(chan bool)
	go howdyGopher()
	mutex.Lock()
	greet="Hello"
	fmt.Println(greet)
	mutex.Unlock()
	<-howdyDone
}


func howdyGopher(){
	mutex.Lock()
	greet="GREET"
	mutex.Unlock()
	howdyDone<- true
}