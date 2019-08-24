package client

import (
	log "github.com/sirupsen/logrus"
	"time"
)

const url = "http://localhost:8080/test"

func StartEchoClient() {

	for {
		client := newHttpClient()
		res := client.Get(url)

		log.Print("Client receive with code ", res.GetStatusCode(), " status ", res.GetStatus())

		time.Sleep(1 * time.Second)
	}

}
