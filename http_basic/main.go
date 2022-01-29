package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("https://www.google.com")
	if err != nil {
		log.Fatalf("Encountered ", err)

	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}
