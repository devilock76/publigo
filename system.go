package main

import (
	"os"
	"log"
	"io/ioutil"
	"encoding/json"
)

type SysSettings struct {
	RootPath	string	`json:"RootPath"`
	TemplPath	string	`json:"TemplPath"`
	DB			string	`json:"DB"`
	Port		string	`json:"Port"`
	ServeStatic	bool	`json:"ServeStatic"`
}

var CurrentConfig SysSettings

func GetSettings (configFile string) {
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	//config = SysSettings{}
	json.Unmarshal(file, &CurrentConfig)
}
