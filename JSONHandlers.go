package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"log"
)

func GetJSONPosts (w http.ResponseWriter, r *http.Request) {
	somePosts := CreateSamplePosts()
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Content-Type", "text/json; charset=utf-8")
    w.WriteHeader(http.StatusOK)
	err := enc.Encode(somePosts)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Data requested")
		log.Println(somePosts)
	}
/*	for _, somePost := range somePosts() {
		err := enc.Encode(somePost)
		if err != nil {
			log.Println(err)
		}
	}*/
}

func CreateSamplePosts () JSONPosts {
	posts := JSONPosts{}
	for i := 0; i < 5; i++ {
		title := "Post " + strconv.Itoa(i+1)
		content := "This is content for post " + strconv.Itoa(i+1)
		post := JSONPost{Title: title, Content: content}
		posts = append(posts, post)
	}
	return posts
}
