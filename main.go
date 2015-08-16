package main

import (
	"log"
	"fmt"
	"net/http"
	"flag"
	"github.com/justinas/alice"
)

func main() {
	
	configPtr := flag.String("path", ".", "Default is PWD")
	flag.Parse()
	
	CurrentConfig := GetSettings(fmt.Sprint(*configPtr, "/config.json"))
	
	log.Println(CurrentConfig)
	
	router := NewRouter()
	baseChain := alice.New(LogHandler, ErrorHandler).Then(router)
	
	fmt.Printf("Listening on port %s\n", fmt.Sprint(":", CurrentConfig.Port))
	
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", CurrentConfig.Port), baseChain))
}
