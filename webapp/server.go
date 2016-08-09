/*
Package webAppExample implements a simple web application in go which implements a simple
TODO list service
API :
	- POST /todo
		- Create a new TODO item
	- DELETE /todo/{id}
		- Deletes an existing TODO item
	- GET /todo/{id}
		- Retrieves an existing TODO item
	- GET /todo
		- Lists all TODO items (metadata only - id, title, due date, status - no description)
		- Optional query params:
			- duedate={date} - list only TODO items with the given due date (default is all)
			- status={new|overdue} list only TODO items with the given status (default is all)
 */
package webapp

import (
	"net/http"
	"fmt"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

var Todo_map map[int]TodoItem

/*
 Handler functions :
 */

// handler function of get request for localhost:8080/
// only prints a welcome message
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to TODO WEB app!")
}

// handler for post request /todo , creates a n todo item and adds it to the TO_do map.
// The input comes in JSON format.
// also handles GET request /todo , outputs all todo items with optional query requests.
// if query is date , assume the input string is in the format: year/month/day , and it will
// be converted to a Date type represnting the start of the day .
func PostNewTodoItemHandlerOrGetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodGet  {
		http.Error(w, "Expected post or get request", 400)
		return
	}
	if r.Method == http.MethodPost {
		t := new( TodoItem )
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(t)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		_, ok := Todo_map[t.Id]
		if ok == false {
			Todo_map[t.Id] = *t
		} else {
			http.Error(w, "Id already exists", 400)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		url := "http://localhost:8080/todo/" + strconv.Itoa(t.Id)
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(*t)

	} else {
		dueDate := r.URL.Query().Get("dueDate")
		status := r.URL.Query().Get("status")
		if status == "" && dueDate == "" {
			for _,v := range Todo_map {
				json.NewEncoder(w).Encode(v)
			}
		} else if status != "" {
			for k,_ := range Todo_map {
				st, _ := strconv.Atoi(status)

				if Todo_map[k].Status == TodoStatus(st) {
					json.NewEncoder(w).Encode(Todo_map[k])
				}
			}
		} else {
			for k,_ := range Todo_map {
				date_input := strings.Split(dueDate, "/")
				year , _ := strconv.Atoi(date_input[0])
				month , _ := strconv.Atoi(date_input[1])
				day , _ := strconv.Atoi(date_input[2])

				if Todo_map[k].DueDate.After(time.Date(year,time.Month(month),day,0,0,0,0,time.UTC)) {
					json.NewEncoder(w).Encode(Todo_map[k])
				}
			}
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

	}

}

// without JSON:
func DeleteAndGetTodoItem(w http.ResponseWriter, r *http.Request){
	fmt.Println("starting del or get req")
	if r.Method != http.MethodDelete && r.Method != http.MethodGet  {
		http.Error(w, "Expected Delete or get request", 400)
		return
	}
	id , _ := strconv.Atoi(r.URL.Path[len("/todo/"):])
	t, ok := Todo_map[id]
	if r.Method == http.MethodDelete {
		if ok == true {
			delete(Todo_map , id )
		} else {
			http.Error(w, "id doesn't exists", 400)
		return
		}
	} else {
		if ok == true {
			json.NewEncoder(w).Encode(t)
		} else {
			http.Error(w, "id doesn't exists", 400)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}












