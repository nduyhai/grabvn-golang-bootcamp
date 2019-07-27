package bootcamp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
)

var globalDB *gorm.DB

func Handle() {
	//load config
	var config Conf
	config.getConf()

	args := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.DB.Host, config.DB.Port, config.DB.User, config.DB.DBName, config.DB.Password)
	db, err := gorm.Open("postgres", args)
	defer db.Close()

	if err == nil {

		err = db.AutoMigrate(Todo{}).Error
		if err != nil {
			log.Fatal("failed to migrate table todos")
		}
		globalDB = db

		r := gin.Default()

		r.GET("/api/todo", getAllTodo)
		r.GET("/api/todo/:id", getTodoById)
		r.POST("/api/todo", createTodo)

		_ = r.Run(":" + config.Server.Port)

	} else {
		log.Fatal("Cannot connect DB: " + err.Error())
	}
}

type Todo struct {
	gorm.Model
	Title     string
	Completed bool
}

func createTodo(c *gin.Context) {
	var argument struct {
		Title string
	}

	err := c.BindJSON(&argument)
	if err != nil {
		c.String(400, "invalid param")
		return
	}

	todo := Todo{
		Title:     argument.Title,
		Completed: false,
	}

	err = globalDB.Create(&todo).Error
	if err != nil {
		c.String(500, "failed to create new todo")
	} else {
		c.JSON(200, todo)
	}
}

func getAllTodo(c *gin.Context) {
	var todo []Todo
	err := globalDB.Find(&todo).Error

	if err != nil {
		c.String(500, "failed to list todolist")
	} else {
		c.JSON(200, todo)
	}

}

func getTodoById(ctx *gin.Context) {
	param := ctx.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err == nil {
		var todo Todo
		globalDB.Where("id = ?", id).Find(&todo)

		ctx.JSON(200, todo)
	} else {
		log.Print(err.Error())
		ctx.String(500, "failed to get todo")
	}
}
