package main

import (
	"fmt"
 
)

func main() {

	var list[]string=make([]string, 2)
	fmt.Println(list)
	fmt.Println(append(list, "ABC"))
	list=append(list, "ABC");
 	fmt.Println(list)
 	fmt.Println("----------")
 	fmt.Println(list[len(list)-1])
 	fmt.Println(len(list))
 	fmt.Println(cap(list))
 	fmt.Println("----------")
	 list=append(list, "ABC");
	 fmt.Println(len(list))
 	fmt.Println(cap(list))
 	fmt.Println("----------")
 	 list=append(list, "ABC");
	 fmt.Println(len(list))
 	fmt.Println(cap(list))
	 list[0]="A"
	 list[1]="B"
	 list[2]="C"
	 list[3]="D"
	 fmt.Println(list)
	 s:=list[1:3]
	 fmt.Println(s)
	 s=list[:3]
	 fmt.Println(s)
	 	 s=list[1:]
	 fmt.Println(s)

	 fmx:=[]string{"A","B","C"}

	 fmt.Println(fmx)
	 fmt.Println(append(fmx, "ABC"))

}