package main

import (
	"log"
	"fmt"
	"net/http"
)

func main() {
	
	CurrentConfig := GetSettings("./config.json")
	
	log.Println(CurrentConfig)

	router := NewRouter()
	
	fmt.Printf("Listening on port %s\n", CurrentConfig.Port)

	log.Fatal(http.ListenAndServe(CurrentConfig.Port, router))
}
