package main

import (
	"log"

	"github.com/matty-rose/distlog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
