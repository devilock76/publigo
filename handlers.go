package main

import (
	"fmt"
	"net/http"
	"html/template"
	"strconv"
	//"github.com/justinas/alice"
)

// Handler for basic logging
//func loggingHandler(h http.Handler) http.Handler {
//	return http.LoggingHandler(h, 1*time.Second, "timed out")
//}

// Handler for user authorization
//func authHandler(h http.Handler) http.Handler {
//	return http.AuthHandler(h, user, "Authenticated")
//}

var templates = template.Must(template.ParseFiles("templates/main.html", "templates/post.html"))

func GetIndex(w http.ResponseWriter, r *http.Request) {
	title := "Welcome"
	content := `This is the welcome page for this go document<br>
	There was a new line <br>
	<br>
	There was 2`
	
	indexPosts := GetRecentPosts()

	p := &Page{Title: title, Content: template.HTML(content), PagePosts: indexPosts}

    err := templates.ExecuteTemplate(w, "main", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetCategoryIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome! Category Index")
}

func GetPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome! Page")
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome! Post")
}

func GetRecentPosts() []Post {
	recent := []Post{}
	for i := 0; i < 4; i++ {
		title := "Post " + strconv.Itoa(i+1)
		content := "This is content for post " + strconv.Itoa(i+1)
		post := Post{Title: title, Content: template.HTML(content)}
		recent = append(recent, post)
	}
	return recent
}
