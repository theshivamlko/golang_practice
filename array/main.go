package main

import (
	"fmt"
 
)

func main() {

	var array [5]int

	fmt.Println(array)

	array[1] = 299

	fmt.Println(array)
	fmt.Println(len(array))
 	


	 var names [5]string 
	  names2:= [2]string{"Shivam","Shubham"}
	fmt.Println(names)
	fmt.Println(names2)

	names3:=names2
	names2[1]="ABC"
	fmt.Println(names3)
	fmt.Println(names2)

	if(names3[0]==names2[0]){
		fmt.Println("MATCH")

	}
	if(names3[0]=="shivam"){
		fmt.Println("MATCH")

	}

}
