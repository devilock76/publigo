package main

import (
	"html/template"
)

type Page struct {
    Title		string			//`json:"title"`
    Content		template.HTML	//`json:"content"`
    PagePosts	[]Post
}

type Pages []Page

type Post struct {
    Title	string			//`json:"title"`
    Content	template.HTML	//`json:"content"`
}

type Posts []Post
