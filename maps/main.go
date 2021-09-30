package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("------------")
	
	var caps map[string]string=make(map[string]string)
	fmt.Println(caps)
	caps["country"]="India"
	fmt.Println(caps)
	fmt.Println(caps["country"])

	for i := 1; i <= 10; i++ {
		caps[strconv.Itoa(i)]=strconv.Itoa(i*10)
	}
for i := 0; i < 10; i++ {
	fmt.Println(caps[strconv.Itoa(i)])
	}

}