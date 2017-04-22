package main

import (
	"fmt"
	"net/http"
)

func NewentryHandler(w http.ResponseWriter, r *http.Request) {

	t := "login: false"

	if checkLogin(r) == true {

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
