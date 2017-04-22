package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

type lista struct {
	Login        bool
	Articletitle []template.HTML
	LoginText    string
	Inventarlist []string
	ArticleText  []template.HTML
}

const (
	// See http://golang.org/pkg/time/#Parse
	timeFormat = "2006-01-02 15:04 MST"
)

var templatesIndex = template.Must(template.ParseFiles("index.html"))
var users string
var name []template.HTML
var ArticleTitle []template.HTML
var timecache int64 = time.Now().Unix()

func IndexHandler(w http.ResponseWriter, r *http.Request) { // Das ist der IndexHandler
	login := false

	if int64(time.Now().Unix()) > timecache+5 {
		cache()
	}

	t := "login: false"
	if checkLogin(r) == true {
		t = "login: true"
		login = true
	}

	lists := lista{login, ArticleTitle, t, readfiles(), name}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

	} else {
		templatesIndex.Execute(w, lists)
	}
}

func cache() {
	name = name[:0]
	ArticleTitle = ArticleTitle[:0]

	//nameTmp := r.FormValue("Name")

	ids, err := db.Query("SELECT id, title, LEFT (text,200) FROM `article` ORDER BY id DESC LIMIT 5")
	checkErr(err)
	fmt.Println(ids)

	for ids.Next() {
		var id int
		var title string
		var text string
		_ = ids.Scan(&id, &title, &text)
		checkErr(err)
		name = append(name, template.HTML(bluemonday.UGCPolicy().SanitizeBytes(blackfriday.MarkdownCommon([]byte(text)))))
		ArticleTitle = append(ArticleTitle, template.HTML(bluemonday.UGCPolicy().SanitizeBytes(blackfriday.MarkdownCommon([]byte(title)))))
		fmt.Println(id, title, text)
	}

	timecache = time.Now().Unix()
}
