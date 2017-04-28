package main

import (
	"html/template"
	"net/http"
	"net/url"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

type view struct {
	Path  string
	Title template.HTML
	Text  template.HTML
}

var templatesView = template.Must(template.ParseFiles("view.html"))

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.Path)
	encodetpath1 := strings.Split(u.Path, "/")

	ids, err := db.Query("SELECT id,title,text FROM article where title=(?)", encodetpath1[2])
	checkErr(err)

	ids.Next()
	var id int
	var title string
	var text string
	_ = ids.Scan(&id, &title, &text)
	checkErr(err)

	TitleTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes(blackfriday.MarkdownCommon([]byte(title))))
	TextTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes(blackfriday.MarkdownCommon([]byte(text))))

	views := view{encodetpath1[2], TitleTMP, TextTMP}
	templatesView.Execute(w, views)

}
