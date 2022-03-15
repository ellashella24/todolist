package todolist

import (
	formatter "todolist-be/delivery/formatter/todolist"
	"todolist-be/delivery/service/todolist"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodolistController struct {
	todolistService todolist.TodolistService
}

func NewTodolistController(todolistService todolist.TodolistService) *TodolistController {
	return &TodolistController{todolistService}
}

func (tc *TodolistController) GetAll(c *gin.Context) {
	todolists, err := tc.todolistService.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": todolists,
	})
}

func (tc *TodolistController) Create(c *gin.Context) {
	var input formatter.InputTodolist

	c.ShouldBind(&input)

	newTodolist, err := tc.todolistService.Create(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": newTodolist,
	})
}

func (tc *TodolistController) Update(c *gin.Context) {
	var input formatter.InputTodolist

	id, _ := strconv.Atoi(c.Param("id"))

	c.ShouldBind(&input)

	updatedTodolist, err := tc.todolistService.Update(id, input)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": updatedTodolist,
	})
}

func (tc *TodolistController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := tc.todolistService.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": nil,
	})
}
