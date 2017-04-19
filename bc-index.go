package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/russross/blackfriday"
)

type lista struct {
	Login        bool
	Title        string
	Title2       string
	Inventarlist []string
	Title4       []template.HTML
}

const (
	// See http://golang.org/pkg/time/#Parse
	timeFormat = "2006-01-02 15:04 MST"
)

var templatesIndex = template.Must(template.ParseFiles("index.html"))
var users string
var name []template.HTML
var timecache int64 = time.Now().Unix()

func IndexHandler(w http.ResponseWriter, r *http.Request) { // Das ist der IndexHandler
	login := false

	if int64(time.Now().Unix()) > timecache+5 {

		cache()
	}

	cookie, _ := r.Cookie("pwd")
	t := "login: false"
	if cookie.Value == "PASSWORD" {
		t = "login: true"
		login = true
	}
	fmt.Println(login)
	lists := lista{login, r.URL.Path, t, readfiles(), name} //r.URL.Path gibt den URL pfad aus
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

	} else {
		templatesIndex.Execute(w, lists)
	}
}

func cache() {
	name = name[:0]

	//nameTmp := r.FormValue("Name")

	rows, err := db.Query("SELECT dh FROM ( SELECT * FROM kk ORDER BY id DESC LIMIT 5) sub ORDER BY id ASC")
	checkErr(err)

	for rows.Next() {
		var text string
		err = rows.Scan(&text)
		checkErr(err)
		name = append(name, template.HTML(blackfriday.MarkdownCommon([]byte(text))))
	}

	timecache = time.Now().Unix()
}
