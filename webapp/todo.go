package webapp

import "time"

const (
	INVALID_ID int = 1 + iota

)

type errorString struct {
	s string

}

func (e *errorString) Error() string {
	return e.s
}

// Define status type:
type TodoStatus int

// Define status possible values:
const (
	newTask  TodoStatus = 1  + iota
	overdue
)

// Define the todoItem :

type TodoItem struct{
	Id int			`json:"id"`
	Title string		`json:"title"`
	DueDate time.Time	`json:"dueDate"`
	Status TodoStatus	`json:"status"`
}

// constructor :
func NewTodoItem(id int, title string, dueDate time.Time ) (*TodoItem , error) {
	t := new(TodoItem )
	t.Id = id
	//t.Title = title
	//t.DueDate = dueDate
	/*
	if dueDate.After(time.Now()) {
		t.Status = newTask
	} else {
		t.Status = overdue
	}
	*/
	return t , nil
}



