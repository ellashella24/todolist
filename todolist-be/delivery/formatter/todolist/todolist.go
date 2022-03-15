package todolist

type InputTodolist struct {
	Name   string `json:"name" form:"name" uri:"name"`
	Status bool   `json:"status" form:"status" uri:"status"`
}
