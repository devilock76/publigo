package main

import (
	"html/template"
	"time"
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

type JSONPost struct {
    Title	string	`json:"Title"`
    Content	string	`json:"Content"`
}

type JSONPosts []JSONPost

type FlexPost struct {
	ID			int
	Cat			int
	Title		string
	Content		string
	Meta		[]string
	Public		int
	Published	time.Time
	FlexField	[]string
	FlexValue	[]string
}

type FlexPosts []FlexPost

type PubliCat struct {
	ID			int
	Name		string
}

type PubliMenu struct {
	Title		string
	URL			string
	Order		int
	Parent		int
}

type PubliMenus []PubliMenu

type PubliMeta struct {
	MetaText	string
	ItemType	[]int
	Item		[]int
}

type PubliMetas []PubliMeta

type Publibrary struct {
	Path		string
	Name		string
	Type		string
}

type PubliUser struct {
	UserName	string
	FullName	string
	MailAddr	string
	PassHash	string
	PassSalt	string
	UserGroup	string
}

type PubliGroup struct {
	GroupName	string
	GroupDesc	string
	AccessCat	[]int
	AccessPerm	[]int
}

type PubliSess struct {
	SessKey		string
	UserName	string
}
