package main

import (
	"fmt"
	"net/http"
)

func NewentryHandler(w http.ResponseWriter, r *http.Request) {

	cookie, _ := r.Cookie("test")

	t := "login: false"
	if cookie.Value == "123" {
		t = "login: true"

		newT := r.FormValue("Name")

		fmt.Println(newT)

		db.Exec("INSERT INTO kk(dh) VALUES(?)", newT)

		checkErr(err)

		fmt.Println("newT")

		http.ServeFile(w, r, "newentry.html")

	}
	fmt.Fprintf(w, `You have to login to do this! -> %s`, t)

}
