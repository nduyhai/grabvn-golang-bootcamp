package feedbackserver

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grabvn-golang-bootcamp/internal/bootcamp/configuration"
	"grabvn-golang-bootcamp/internal/bootcamp/feedback"
	"log"
	"net"
)

func StartServer() {
	var config configuration.Conf
	config.LoadConf()

	db, err := connectDB(config)
	if err == nil {
		log.Print("begin exec auto migrate....")
		err = db.AutoMigrate(Feedback{}).Error
		if err != nil {
			log.Fatal("failed to migrate table Feedback")
		}

		log.Print("begin init rpc server....")

		ln, err := net.Listen("tcp", ":"+config.RPC.Port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		cred, err := credentials.NewServerTLSFromFile(config.RPC.CertFile, config.RPC.KeyFile)
		if err != nil {
			log.Fatalf("could not load TLS keys:: %v", err)
		}

		rpcServer := grpc.NewServer(grpc.Creds(cred))
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

func connectDB(config configuration.Conf) (*gorm.DB, error) {
	args := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.DB.Host, config.DB.Port, config.DB.User, config.DB.DBName, config.DB.Password)
	db, err := gorm.Open("postgres", args)
	return db, err
}
