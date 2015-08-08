package main

type Page struct {
    Title	string	`json:"title"`
    Content	bool	`json:"content"`
}

type Pages []Page

type Post struct {
    Title	string	`json:"title"`
    Content	bool	`json:"content"`
}

type Posts []Post
