package bootcamp

import (
	"github.com/gin-gonic/gin"
	"log"
)

const Port = "8080"

func StartServer() {
	log.Print("begin setup http server...")

	webContext := NewWebContext()
	server := initializeServer()
	setupRoute(server, webContext)

	log.Print("begin run http server...")

	_ = server.Run(":" + Port)
}
func initializeServer() *gin.Engine {
	server := gin.New()

	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	return server
}
func setupRoute(server *gin.Engine, webContext *WebContext) {
	server.GET("/api/query", webContext.QueryUser)

}
