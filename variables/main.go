package main

import "fmt"

func main() {
	fmt.Println("variables")
	var i int
	fmt.Println(i)
	var j,k int =10,24
	fmt.Println(j,k)
	m,n:=4.4,6
	fmt.Println(m,n)
	fmt.Println("m: %d n:",m,n)
	const pi float32=3.14
	fmt.Println(pi)

	var x bool=true
	fmt.Println(x)

	const(
		light=300000
		randNo float32=102.2
	)
	fmt.Println(light,randNo)

	var name string ="Shivam"
	fmt.Println(string("Shubham"[3]))
	fmt.Println(name[0])
	fmt.Println(string(name[0]))
	fmt.Println("Hello " +name+"!")

}
