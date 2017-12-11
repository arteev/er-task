package main

import (
	"log"

	"os"

	"github.com/arteev/er-task/src/app"
)

func main() {
	addr := ":8080"
	if eaddr, ok := os.LookupEnv("APPADDR"); ok {
		addr = eaddr
	}
	connection := "postgres://postgres:example@127.0.0.1/carrental?sslmode=disable"
	if conn, ok := os.LookupEnv("POSTGRES"); ok {
		connection = conn
	}
	rconn := ""
	if redis, ok := os.LookupEnv("REDIS"); ok {
		rconn = redis
	}
	if err := app.Run(addr, connection, rconn); err != nil {
		log.Fatal(err)
	}
}
