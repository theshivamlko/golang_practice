package custompackage

import "fmt"

func PrintGreetings(){
	fmt.Println("PrintGreetings")
}

// optinal , first to run
func  init()  {
	fmt.Println("Greetings INIT")
 
}
