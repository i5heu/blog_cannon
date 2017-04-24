package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

type Article struct {
	Articletitle template.HTML
	ArticleText  template.HTML
}

type lista struct {
	Login        bool
	LoginText    string
	Inventarlist []string
	Articles     []Article
}

const (
	// See http://golang.org/pkg/time/#Parse
	timeFormat = "2006-01-02 15:04 MST"
)

var templatesIndex = template.Must(template.ParseFiles("index.html"))
var timecache int64 = time.Now().Unix()
var tmp []Article

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

	lists := lista{login, t, readfiles(), tmp}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

	} else {
		templatesIndex.Execute(w, lists)
	}
}

func cache() {
	tmp = tmp[:0]

	ids, err := db.Query("SELECT id, title, LEFT (text,200) FROM `article` ORDER BY id DESC LIMIT 5")
	checkErr(err)
	fmt.Println(ids)

	for ids.Next() {
		var id int
		var title string
		var text string
		_ = ids.Scan(&id, &title, &text)
		checkErr(err)

		TitleTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes(blackfriday.MarkdownCommon([]byte(title))))
		TextTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes(blackfriday.MarkdownCommon([]byte(text))))

		tmp = append(tmp, Article{TitleTMP, TextTMP})

		fmt.Println(id, TitleTMP, TextTMP)
	}

	timecache = time.Now().Unix()
}
