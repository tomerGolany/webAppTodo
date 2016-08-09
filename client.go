package main

import (
	"net/http"
	"webAppExample/webapp"
	"fmt"
)

var (
	Version string
	Build  string
)


func main(){
	fmt.Println("Version:", Version)
	fmt.Println("Build:", Build)
	webapp.Todo_map = make(map[int]webapp.TodoItem)
	//http.HandleFunc, tells the http package to handle all requests to the web
	//root ("/") with handler.
	http.HandleFunc("/", webapp.Index)
	http.HandleFunc("/todo", webapp.PostNewTodoItemHandlerOrGetAll)
	http.HandleFunc("/todo/", webapp.DeleteAndGetTodoItem)
	//http.ListenAndServe, specifying it should listen on port 8080 on any interface (":8080").
	// (Don't worry about its second parameter, nil, for now.) This function will block until the program is terminated.
	http.ListenAndServe(":8080", nil)
}
