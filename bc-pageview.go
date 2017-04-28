package main

import (
	"html/template"
	"net/http"

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
	path := r.URL.Path[3:]

	ids, err := db.Query("SELECT id,title,text FROM article where id=(?)", path)
	checkErr(err)

	ids.Next()
	var id int
	var title string
	var text string
	_ = ids.Scan(&id, &title, &text)
	checkErr(err)

	TitleTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes(blackfriday.MarkdownCommon([]byte(title))))
	TextTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes(blackfriday.MarkdownCommon([]byte(text))))

	views := view{path, TitleTMP, TextTMP}
	templatesView.Execute(w, views)

}
