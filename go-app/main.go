package main

import (
	"go-app/server"
	"log"
)

func main() {
	err := server.NewServer().StartServer(3000)
	if err != nil {
		log.Fatal(err)
	}
}
