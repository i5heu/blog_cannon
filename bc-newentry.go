package main

import (
	"fmt"
	"net/http"
)

func NewentryHandler(w http.ResponseWriter, r *http.Request) {

	t := "login: false"

	if checkLogin(r) == true {
		//if true == true {
		t = "login: true"

		newT := r.FormValue("Name")
		newTitle := r.FormValue("Title")

		db.Exec("INSERT INTO article(title,text) VALUES(?,?)", newTitle, newT)

		checkErr(err)

		http.ServeFile(w, r, "newentry.html")

	}
	fmt.Fprintf(w, `You have to login to do this! -> %s`, t)

}
