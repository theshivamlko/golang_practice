package main

import (
	"html/template"
	"log"
	"net/http"
)


type Gopher struct{
	Name string
}

func main() {
	
http.HandleFunc("/username",username)
http.ListenAndServe(":8081",nil)


}

func username(w http.ResponseWriter, r *http.Request){
	user:=r.URL.Query().Get("user")

	gopher:=Gopher{Name:user}
	t, err:=template.ParseFiles("greeting.html")

	if err!=nil{
		log.Fatal("Error ",err)
	}

	t.Execute(w,gopher)

	if user==""{
		
		//http.Error(w,"Username",http.StatusBadRequest)
	}else{
		// fmt.Fprint(w,"Username ",user,http.StatusAccepted,userResult)

	}


 
}