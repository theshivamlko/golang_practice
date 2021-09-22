package main

import (
	"fmt"
	"os"
	"log"
	"strconv"
	"flag"
)

func main() {

	var repeat int

	var err error

	if len(os.Args)>=2 {
		repeat,err=strconv.Atoi(os.Args[1])
		fmt.Println(os.Args[0])

		if err!=nil{
			log.Fatal(err)
		}

		for i:=0;i<repeat ; i++ {
			fmt.Println("Hello ",i)
		}
	} else {
		fmt.Println("No input ")
	}

	var gopherName string
	flag.StringVar(&gopherName,"gopher","Gopher","Gopher2")
	flag.Parse()
	fmt.Println("hello "+ gopherName + "!")


	
}
