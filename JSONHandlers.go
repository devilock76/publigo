package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"log"
)

func GetJSONPosts (w http.ResponseWriter, r *http.Request) {
	somePosts := CreateSamplePosts
	enc := json.NewEncoder(w)
	for _, somePost := range somePosts() {
		err := enc.Encode(somePost)
		if err != nil {
			log.Println(err)
		}
	}
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