package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"time"
)

type User struct {
	id       int
	name     string
	password string
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func Db_connect() *sql.DB {
	db, err := sql.Open("mysql", "go_user:qwe123@/DnD_Users") //Func_db
	if err != nil {
		panic(err)
	}
	return db
}

func NewUser(db *sql.DB, name string, pas string) {
	rand.Seed(time.Now().UTC().UnixNano())

	result, err := db.Exec("insert into DnD_Users.Users (id, username, password) values (?, ?, ?)",
		randInt(1, 100000), name, pas)
	fmt.Println("Запись в БД прошла успешно", result)
	if err != nil {
		panic(err)
	}
}

func FindUser(db *sql.DB, name string, pas string) (User, bool) {
	var user User
	row := db.QueryRow("SELECT id, username, password FROM DnD_Users.Users WHERE username = ? AND password = ?", name, pas)
	err := row.Scan(&user.id, &user.name, &user.password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, false
		}
	}
	return user, true
}

//func Authentication(db *sql.DB, name string, pas string) (User, bool) {
//
//	user, exictence := FindUser(db, name, pas)
//	if exictence == true {
//		return user, true
//	}
//	return user, false
//}
