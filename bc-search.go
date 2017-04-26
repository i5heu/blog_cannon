package main

import (
	"html/template"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

type SearchResult struct {
	ArticleId    int
	Articletitle template.HTML
	ArticleText  template.HTML
}

type search struct {
	Searchterm   string
	SeachResults []SearchResult
}

var templatesSearch = template.Must(template.ParseFiles("search.html"))
var tmpSearch []SearchResult

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	tmpSearch = tmpSearch[:0]
	searchterm := r.URL.Path[3:]

	newquery := "*" + searchterm + "*"
	ids, err := db.Query("SELECT  id,title,SUBSTR(text,1,100) FROM article WHERE MATCH (title,text) AGAINST (? IN BOOLEAN MODE)", newquery)
	checkErr(err)
	for ids.Next() {
		var id int
		var title string
		var text string
		_ = ids.Scan(&id, &title, &text)
		checkErr(err)
		text = text + "..."

		TitleTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes(blackfriday.MarkdownCommon([]byte(title))))
		TextTMP := template.HTML(bluemonday.UGCPolicy().SanitizeBytes(blackfriday.MarkdownCommon([]byte(text))))
		tmpSearch = append(tmpSearch, SearchResult{id, TitleTMP, TextTMP})
	}
	searchs := search{searchterm, tmpSearch}

	templatesSearch.Execute(w, searchs)

}
