package feedbackclient

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"grabvn-golang-bootcamp/internal/bootcamp/configuration"
	"grabvn-golang-bootcamp/internal/bootcamp/feedback"
	"log"
)

func StartClient() {
	var config configuration.Conf
	config.LoadConf()

	log.Print("begin init rpc client....")

	conn, err := grpc.Dial("localhost:"+config.RPC.Port, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("failed to connect server: %v", err)
	}
	defer conn.Close()
	client := feedback.NewFeedbackServiceClient(conn)

	webContext := &WebProxy{client: &client}
	server := initializeServer()
	setupRoute(server, webContext)

	log.Print("begin run http server...")

	_ = server.Run(":" + config.Server.Port)
}

func initializeServer() *gin.Engine {
	server := gin.New()

	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	return server
}

func setupRoute(server *gin.Engine, webContext *WebProxy) {
	server.POST("/api/feedback", webContext.addFeedback)
	server.GET("/api/feedback/:id", webContext.getFeedbackById)
	server.DELETE("/api/feedback/:id", webContext.deleteFeedbackById)
	server.GET("/api/code/:id", webContext.getFeedbackByCode)
}
