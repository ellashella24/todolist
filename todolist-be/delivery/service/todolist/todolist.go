package todolist

import (
	formatter "todolist-be/delivery/formatter/todolist"
	"todolist-be/entity"
	"todolist-be/repository/todolist"
)

type TodolistService interface {
	GetAll() ([]entity.Todolist, error)
	Create(input formatter.InputTodolist) (entity.Todolist, error)
	Update(id int, input formatter.InputTodolist) (entity.Todolist, error)
	Delete(id int) (entity.Todolist, error)
}

type todolistService struct {
	todolistRepo todolist.TodolistRepository
}

func NewTodolistService(todolistRepo todolist.TodolistRepository) *todolistService {
	return &todolistService{todolistRepo}
}

func (us *todolistService) GetAll() ([]entity.Todolist, error) {
	todolists, err := us.todolistRepo.GetAll()

	if err != nil {
		return nil, err
	}

	return todolists, nil
}

func (ts *todolistService) Create(input formatter.InputTodolist) (entity.Todolist, error) {
	todolist := entity.Todolist{}

	todolist.Name = input.Name
	todolist.Status = input.Status

	newTodolist, err := ts.todolistRepo.Create(todolist)

	if err != nil {
		return entity.Todolist{}, err
	}

	return newTodolist, err
}

func (ts *todolistService) Update(id int, input formatter.InputTodolist) (entity.Todolist, error) {
	todolist, err := ts.todolistRepo.Find(id)

	if err != nil {
		return todolist, err
	}

	todolist.Name = input.Name
	todolist.Status = input.Status

	updatedTodolist, err := ts.todolistRepo.Update(id, todolist)

	if err != nil {
		return entity.Todolist{}, err
	}

	return updatedTodolist, err
}

func (ts *todolistService) Delete(id int) (entity.Todolist, error) {
	deletedTodolist, err := ts.todolistRepo.Delete(id)

	if err != nil {
		return entity.Todolist{}, err
	}

	return deletedTodolist, err
}
