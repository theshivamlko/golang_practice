package main

import (
	"fmt"
	"net/http"
)

func main() {
	
http.HandleFunc("/hello",hello)
http.HandleFunc("/username",username)
http.ListenAndServe(":8081",nil)


}

func hello(w http.ResponseWriter, r *http.Request){
	fmt.Print("Hello")
}

func username(w http.ResponseWriter, r *http.Request){
	var userResult bool
	user:=r.URL.Query().Get("username")

	if user==""{
		http.Error(w,"Username",http.StatusBadRequest)
	}else{
		fmt.Fprint(w,"Username ",user,http.StatusAccepted,userResult)

	}
 
}