package main

import (
	"log"

	"github.com/arteev/er-task/app"
)

func main() {
	addr := "127.0.0.1:81"
	connection := "postgres://postgres:example@192.168.1.43/carrental?sslmode=disable"

	if err := new(app.App).Run(addr, connection); err != nil {
		log.Fatal(err)
	}
}
