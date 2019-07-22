package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"grabvn-golang-bootcamp/internal/bootcamp"
)

func main() {
	bootcamp.Handle()
}
