package bootcamp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
)

func Handle() {
	//load config
	var config Conf
	config.getConf()

	db, err := connectDB(config)
	if err == nil {

		defer db.Close()

		err = db.AutoMigrate(Todo{}).Error
		if err != nil {
			log.Fatal("failed to migrate table todo")
		}

		webContext := &WebContext{DB: db}
		server := gin.Default()

		server.GET("/api/todo", webContext.getAllTodo)
		server.GET("/api/todo/:id", webContext.getTodoById)
		server.POST("/api/todo", webContext.createTodo)

		_ = server.Run(":" + config.Server.Port)

	} else {
		log.Fatal("Cannot connect DB: " + err.Error())
	}
}

func connectDB(config Conf) (*gorm.DB, error) {
	args := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.DB.Host, config.DB.Port, config.DB.User, config.DB.DBName, config.DB.Password)
	db, err := gorm.Open("postgres", args)
	return db, err
}
