package feedbackserver

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"grabvn-golang-bootcamp/internal/bootcamp/feedback"
	"log"
	"net"
)

const (
	port = ":9000"
)

func StartServer() {
	var config Conf
	config.loadConf()

	db, err := connectDB(config)
	if err == nil {
		err = db.AutoMigrate(Feedback{}).Error
		if err != nil {
			log.Fatal("failed to migrate table Feedback")
		}

		log.Print("begin init rpc server....")

		ln, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		rpcServer := grpc.NewServer()
		feedback.RegisterFeedbackServiceServer(rpcServer, &server{DB: db})
		if rpcServer == nil {
			log.Fatalf("failed to register server: %v", err)
		}
		if err := rpcServer.Serve(ln); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	} else {
		log.Fatal("Cannot connect DB: " + err.Error())
	}
}

func connectDB(config Conf) (*gorm.DB, error) {
	args := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.DB.Host, config.DB.Port, config.DB.User, config.DB.DBName, config.DB.Password)
	db, err := gorm.Open("postgres", args)
	return db, err
}
