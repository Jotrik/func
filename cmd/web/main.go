package main

import (
	"database/sql"
	"log"
	"net/http"
)

var db *sql.DB

func main() {

	//addr := flag.String("addr", ":4000", "Сетевой адрес HTTP")
	//flag.Parse()

	db = Db_connect()
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/auth", auth)
	mux.HandleFunc("/ws", WebsocketHandler)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("http://127.0.0.1:4000\n")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
