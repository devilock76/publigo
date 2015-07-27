package publigo

import (
	"net/http"
)

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
