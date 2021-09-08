package main

import (
	"custompackage"
	f "fmt"
)
 
func main() {
 //   calculator.Calc()

	const (i=10)

	const(
		_=iota*i
		red
		green
		yellow

	)
 	custompackage.PrintGreetings()

	f.Println(red)
	f.Println(green)
	f.Println(yellow)


	
}
