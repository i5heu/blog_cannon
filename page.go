package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"
)

type view struct {
	Title template.HTML
}

var namespaceView = template.Must(template.ParseFiles("./template/namespace.html", HtmlStructHeader, HtmlStructFooter))
var templatesView = template.Must(template.ParseFiles("./template/view.html", HtmlStructHeader, HtmlStructFooter))

var PAGECACHE = make(map[string]template.HTML)
var PAGECACHECACHE = make(map[string]template.HTML)
var PAGECACHEWRITE bool = false
var PAGECACHECACHEWRITE bool = false

func PageHandler(w http.ResponseWriter, r *http.Request) {

	u, err := url.Parse(r.URL.Path)

	checkErr(err)
	encodetpath1 := strings.Split(u.Path, "/")

	if len(encodetpath1) == 3 {
		CategoryHandler(w, r)
		return
	}
	if len(encodetpath1) < 4 {
		fmt.Fprintf(w, "ERROR 404")
		return
	}

	views := view{}
	templatesView.Execute(w, views)

}

func CategoryHandler(w http.ResponseWriter, r *http.Request) {

}
