package todolist

import (
	"todolist-be/entity"

	"gorm.io/gorm"
)

type TodolistRepository interface {
	GetAll() ([]entity.Todolist, error)
	Create(newTodolist entity.Todolist) (entity.Todolist, error)
	Update(id int, updatedTodolist entity.Todolist) (entity.Todolist, error)
	Delete(id int) (entity.Todolist, error)
	Find(id int) (entity.Todolist, error)
}

type todolistRepository struct {
	db *gorm.DB
}

func NewTodolistRepository(db *gorm.DB) *todolistRepository {
	return &todolistRepository{db}
}

func (tr *todolistRepository) GetAll() ([]entity.Todolist, error) {
	var todolists []entity.Todolist

	err := tr.db.Find(&todolists).Error
	if err != nil {
		return nil, err
	}

	return todolists, nil
}

func (tr *todolistRepository) Create(newTodolist entity.Todolist) (entity.Todolist, error) {
	err := tr.db.Create(&newTodolist).Error

	if err != nil {
		return entity.Todolist{}, err
	}

	return newTodolist, nil
}

func (tr *todolistRepository) Update(id int, updatedTodolist entity.Todolist) (entity.Todolist, error) {
	tr.db.Save(&updatedTodolist)

	return updatedTodolist, nil
}

func (tr *todolistRepository) Find(id int) (entity.Todolist, error) {
	var todolist entity.Todolist

	err := tr.db.Where("id = ?", id).Find(&todolist).Error

	if err != nil {
		return todolist, err
	}

	return todolist, err
}

func (tr *todolistRepository) Delete(id int) (entity.Todolist, error) {
	var todolist entity.Todolist

	err := tr.db.Where("id = ?", id).Find(&todolist).Error

	if err != nil {
		return entity.Todolist{}, err
	}

	tr.db.Delete(&todolist)

	return todolist, nil
}
