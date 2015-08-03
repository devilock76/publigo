package publigo

import (
	"net/http"
	"github.com/justinas/alice"
)

// Handler for basic logging
func loggingHandler(h http.Handler) http.Handler {
	return http.LoggingHandler(h, 1*time.Second, "timed out")
}

// Handler for user authorization
func authHandler(h http.Handler) http.Handler {
	return http.AuthHandler(h, user, "Authenticated")
}

func GetIndex(h http.Handler) http.Handler {
	return http.GetIndex(h, index, "Returned")
}

func GetCategoryIndex(h http.Handler) http.Handler {
	return http.GetCategoryIndex(h, categoryindex, "Returned")
}

func GetPage(h http.Handler) http.Handler {
	return http.GetPage(h, page, "Returned")
}

func GetPost(h http.Handler) http.Handler {
	return http.GetPost(h, post, "Returned")
}
