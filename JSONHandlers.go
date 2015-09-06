package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"log"
	//"io/ioutil"
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

//func CreateJSONPost (w http.ResponseWriter, r *http.Request) {
	//newPost := JSONPost{}
	//body, err := ioutil.ReadAll(ioutil.LimitReader(r.Body, 1048576))
	//if err != nil {
		//panic(err)
	//}
	//if err := r.Body.Close(); err != nil {
		//panic(err)
	//}
	//if err := json.Unmarshal(body, &newPost); err != nil {
		//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		//w.WriteHeader(422) // unprocessable entity
		//if err := json.NewEncoder(w).Encode(err); err != nil {
			//panic(err)
		//}
	//}
	////newID, err :_ DBSavePost(newPost)
	////if err !=nil {
	////	panic(err)
	////}
//}

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
