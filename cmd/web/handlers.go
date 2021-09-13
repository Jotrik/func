package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/main.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

//func chat(w http.ResponseWriter, r *http.Request) {
//	tmpl, _ := template.ParseFiles("./ui/html/chat.html")
//	if err := tmpl.Execute(w, nil); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}

func auth(rw http.ResponseWriter, rq *http.Request) {
	if rq.Method == http.MethodPost {
		err := rq.ParseForm()
		if err != nil {
			return
		}
		name := rq.FormValue("name")
		pas := rq.FormValue("pas")
		_, ex := FindUser(db, name, pas)
		if ex == true {
			//Authentication(db, name, pas)
			log.Print("Авторизация прошла успешно")
		} else {
			NewUser(db, name, pas)
			log.Print("Регистрация прошла успешно")
		}
		//http.ServeFile(rw, rq, "./ui/html/main.html")
	}
}
