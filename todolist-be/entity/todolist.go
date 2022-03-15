package entity

import "gorm.io/gorm"

type Todolist struct {
	gorm.Model
	Name   string
	Status bool
}
