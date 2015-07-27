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
