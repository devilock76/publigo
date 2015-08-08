package main

import (
	"html/template"
)

type Page struct {
    Title	string			//`json:"title"`
    Content	template.HTML	//`json:"content"`
}

type Pages []Page

type Post struct {
    Title	string	`json:"title"`
    Content	bool	`json:"content"`
}

type Posts []Post
