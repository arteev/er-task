package main

import (
	"log"

	"os"

	"github.com/arteev/er-task/src/app"
)

func main() {
	addr := ":8082"
	if eaddr, ok := os.LookupEnv("APPADDR"); ok {
		addr = eaddr
	}
	connection := "postgres://postgres:example@127.0.0.1/carrental?sslmode=disable"
	if conn, ok := os.LookupEnv("POSTGRES"); ok {
		connection = conn
	}

	if err := new(app.App).Run(addr, connection); err != nil {
		log.Fatal(err)
	}
}
