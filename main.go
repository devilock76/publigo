package main

import (
	"log"
	"fmt"
	"net/http"
)

func main() {

	router := NewRouter()
	
	fmt.Println("Listening on port 3000")

	log.Fatal(http.ListenAndServe(":3000", router))
}
