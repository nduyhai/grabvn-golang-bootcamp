package feedbackserver

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grabvn-golang-bootcamp/internal/bootcamp/configuration"
	"grabvn-golang-bootcamp/internal/bootcamp/feedback"
	"io/ioutil"
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

		cert, err := loadX509KeyPair(config.RPC.CertFile, config.RPC.KeyFile, config.RPC.Passphrase)
		if err != nil {
			log.Fatalf("could not load TLS keys:: %v", err)
		}
		cred := credentials.NewServerTLSFromCert(&cert)

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

func loadX509KeyPair(certFile, keyFile, passphrase string) (tls.Certificate, error) {
	certPEMBlock, err := ioutil.ReadFile(certFile)
	if err != nil {
		return tls.Certificate{}, err
	}
	keyPEMBlock, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return tls.Certificate{}, err
	}
	block, rest := pem.Decode(keyPEMBlock)
	if len(rest) > 0 {
		return tls.Certificate{}, errors.New("extra data")
	}
	der, err := x509.DecryptPEMBlock(block, []byte(passphrase))
	if err != nil {
		return tls.Certificate{}, err
	}
	if _, err := x509.ParsePKCS1PrivateKey(der); err != nil {
		return tls.Certificate{}, err
	}
	return tls.X509KeyPair(certPEMBlock, der)
}
