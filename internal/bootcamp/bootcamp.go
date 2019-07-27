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
		server := initializeServer()
		setupRoute(server, webContext)

		_ = server.Run(":" + config.Server.Port)

	} else {
		log.Fatal("Cannot connect DB: " + err.Error())
	}
}

func setupRoute(server *gin.Engine, webContext *WebContext) {
	server.GET("/api/todo", webContext.getAllTodo)
	server.GET("/api/todo/:id", webContext.getTodoById)
	server.POST("/api/todo", webContext.createTodo)

	//Add Test Basic auth
	authGroup := server.Group("/auth", gin.BasicAuth(getAllAccount()))

	authGroup.GET("/", func(context *gin.Context) {

		context.String(200, "OK")
	})
}

func initializeServer() *gin.Engine {
	server := gin.New()

	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	return server
}

func connectDB(config Conf) (*gorm.DB, error) {
	args := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.DB.Host, config.DB.Port, config.DB.User, config.DB.DBName, config.DB.Password)
	db, err := gorm.Open("postgres", args)
	return db, err
}
