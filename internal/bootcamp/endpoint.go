package bootcamp

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func createTodo(ctx *gin.Context, webContext *WebContext) {
	var argument struct {
		Title string
	}

	err := ctx.BindJSON(&argument)
	if err != nil {
		ctx.String(400, "invalid param")
		return
	}

	todo := Todo{
		Title:     argument.Title,
		Completed: false,
	}

	err = webContext.DB.Create(&todo).Error
	if err != nil {
		ctx.String(500, "failed to create new todo")
	} else {
		ctx.JSON(200, todo)
	}
}

func getAllTodo(ctx *gin.Context, webContext *WebContext) {
	var todo []Todo
	err := webContext.DB.Find(&todo).Error

	if err != nil {
		ctx.String(500, "failed to list todolist")
	} else {
		ctx.JSON(200, todo)
	}

}

func getTodoById(ctx *gin.Context, webContext *WebContext) {
	param := ctx.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err == nil {
		var todo Todo
		webContext.DB.Where("id = ?", id).Find(&todo)

		ctx.JSON(200, todo)
	} else {
		log.Print(err.Error())
		ctx.String(500, "failed to get todo")
	}
}
