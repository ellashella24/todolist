package route

import (
	"todolist-be/delivery/controller/todolist"

	"github.com/gin-gonic/gin"
)

func RegisterPath(g *gin.Engine, tc *todolist.TodolistController) {
	g.GET("/todolists", tc.GetAll)
	g.POST("/todolist", tc.Create)
	g.PUT("/todolist/:id", tc.Update)
	g.DELETE("/todolist/:id", tc.Delete)
}
