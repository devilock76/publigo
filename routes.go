package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./frontend/"))))

	return router
}

var routes = Routes{
	//Route{
		//"Index",
		//"GET",
		//"/",
		//router.PathPrefix("/frontend/").Handler(http.StripPrefix("/frontend/", http.FileServer(http.Dir("./frontend/")))),
	//},
	Route{
		"Category",
		"GET",
		"/Category/{Category}",
		GetCategoryIndex,
	},
	Route{
		"Page",
		"GET",
		"/Page/{Page}",
		GetPage,
	},
	Route{
		"Post",
		"GET",
		"/Post/{Post}",
		GetPost,
	},
	Route{
		"PostsIndex",
		"GET",
		"/JSON/Posts",
		GetJSONPosts,
	},
}
