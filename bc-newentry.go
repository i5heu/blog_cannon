package main

import (
	"fmt"
	"net/http"
)

func NewentryHandler(w http.ResponseWriter, r *http.Request) {

	var cookie string
	var cookieTMP *http.Cookie

	if sessionExists(r, "pwd") == true {
		cookieTMP, _ = r.Cookie("pwd")
		cookie = cookieTMP.Value
	}

	t := "login: false"
	if cookie == "PASSWORD" {
		t = "login: true"

		newT := r.FormValue("Name")

		fmt.Println(newT)

		db.Exec("INSERT INTO article(text) VALUES(?)", newT)

		checkErr(err)

		fmt.Println("newT")

		http.ServeFile(w, r, "newentry.html")

	}
	fmt.Fprintf(w, `You have to login to do this! -> %s`, t)

}
