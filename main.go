package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mrkaspa/dbadminser/handler"
)

func main() {
	startServer()
}

func startServer() {
	http.Handle("/", handler.NewRouter())
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	err := http.ListenAndServe(host+":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
