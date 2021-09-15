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
	mux.HandleFunc("/authPage", authPage)
	mux.HandleFunc("/ws", WebsocketHandler)

	mux.HandleFunc("/backLink", backLink)
	mux.HandleFunc("/personalArea", personalArea)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/contacts", contacts)
	mux.HandleFunc("/price", price)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	err := http.ListenAndServe(":80", mux) //:80
	log.Fatal(err)
}
