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

	log.Println("http://127.0.0.1:8000\n")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
