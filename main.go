package main

import (
	"log"

	"os"

	"github.com/arteev/er-task/app"
)

func main() {
	addr := ":8082"
	if eaddr, ok := os.LookupEnv("APPADDR"); ok {
		addr = eaddr
	}
	connection := "postgres://postgres:example@0.0.0.0/carrental?sslmode=disable"
	if econn, ok := os.LookupEnv("POSTGRES"); ok {
		connection = econn
	}

	if err := new(app.App).Run(addr, connection); err != nil {
		log.Fatal(err)
	}
}
