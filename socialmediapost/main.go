package main

import (
	"fmt"
	"time"
)

type MoodState int

const(
	_=iota
	MoodHappy
	MoodSad
)

func main() {
	fmt.Println("------------")

	post:=Post{TimeStamp{TimeCreated: time.Now()},"google","google.com",MoodHappy}

	fmt.Println(post)

 

}

type TimeStamp struct{
	TimeCreated time.Time
	TimeModified time.Time
 

}

func NewPost(title string,   URL string,mood MoodState) *Post{
	return &Post{Title: title,URL: URL,Mood: mood}
}

 


type Post struct{
    TimeStamp
	Title string
	URL string
	Mood MoodState

 
}

var Moods map[string]MoodState 

func init(){
  Moods= map[string]MoodState{"happy":MoodHappy,"sad":MoodSad}

}
