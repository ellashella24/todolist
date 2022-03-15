package main

import (
	"todolist-be/config"
	todoController "todolist-be/delivery/controller/todolist"
	"todolist-be/delivery/route"
	todoService "todolist-be/delivery/service/todolist"
	todoRepo "todolist-be/repository/todolist"
	"todolist-be/util"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.GetConfig()
	db := util.InitDB(config)

	g := gin.Default()
	g.Use(cors.Default())

	todolistRepository := todoRepo.NewTodolistRepository(db)
	todolistService := todoService.NewTodolistService(todolistRepository)
	todolistController := todoController.NewTodolistController(todolistService)

	route.RegisterPath(g, todolistController)

	g.Run(":8000")
}
