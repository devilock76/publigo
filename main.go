package publigo

import (
	"fmt"
	"github.com/justinas/alice"
	"net/http"
)

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
