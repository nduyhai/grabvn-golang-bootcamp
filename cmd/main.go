package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

var globalDB *gorm.DB

func main() {
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
		r.POST("/api/todo", createTodo)

		_ = r.Run(":8080")

	} else {
		log.Fatal("Cannot connect DB: " + err.Error())
	}

}

type Todo struct {
	gorm.Model
	Title     string
	Completed bool
}

type Conf struct {
	DB struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	}
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
		return
	}

	c.JSON(200, todo)
}

func getAllTodo(c *gin.Context) {
	var todos []Todo
	err := globalDB.Find(&todos).Error

	if err != nil {
		c.String(500, "failed to list todolist")
		return
	}

	c.JSON(200, todos)
}

func (c *Conf) getConf() *Conf {
	pwd, _ := os.Getwd()
	yamlFile, err := ioutil.ReadFile(pwd + "/configs/application.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
