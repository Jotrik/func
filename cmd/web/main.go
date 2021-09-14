package main

import (
	"database/sql"
	"log"
	"net/http"
)

var db *sql.DB

func main() {

	db = Db_connect()
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/auth", auth)
	mux.HandleFunc("/ws", WebsocketHandler)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("localhost:8000\n")
	err := http.ListenAndServe(":80", mux)
	log.Fatal(err)
}
