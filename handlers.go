package main

import (
	"fmt"
	"net/http"
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

func GetIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome! Index")
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
