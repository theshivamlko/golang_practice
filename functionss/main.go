package main

import "fmt"

var y = 4

func main() {
	odd(y)
	odd(25)
 s,s2 :=	sum(25,50)
	fmt.Println(s,s2)
	s3 :=	multisum(25,50,100)
	fmt.Println(s,s3)

}

func odd(x int) {
	if x%2 == 0 {
		fmt.Println("EVEN")
	} else {
		fmt.Println("ODD")
	}
}


func sum (a int, b int) (s int,s2 int){
	return  a + b, a/b
}
func multisum (args ...int) int{
	s:=0
	for i:=0;i<len(args) ;i++  {
		s=s+args[i]
		
	}
	return  s
}
