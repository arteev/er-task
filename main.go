package main

import (
	"log"

	"github.com/arteev/er-task/app"
)

func main() {
	addr := "127.0.0.1:8082"
	connection := "postgres://postgres:example@127.0.0.1/carrental?sslmode=disable"

	if err := new(app.App).Run(addr, connection); err != nil {
		log.Fatal(err)
	}
}
