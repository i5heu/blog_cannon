package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var templatefolder string = "/home/her/CODE/blog_cannon/template"

var HtmlStructHeader string = templatefolder + `/header.html`
var HtmlStructFooter string = templatefolder + `/footer.html`

var templatesDesktop = template.Must(template.ParseFiles("./template/home.html", HtmlStructHeader, HtmlStructFooter))

var fs = http.FileServer(http.Dir("static"))

type lista struct {
	Rendertime time.Duration
}

func main() {
	fmt.Println("HALLO")

	http.HandleFunc("/favicon.ico", FaviconHandler)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":8080", nil)
}

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/static/favicon/favicon.ico", 302)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	lists := lista{}
	lists = lista{time.Since(start)}

	templatesDesktop.Execute(w, lists)

	fmt.Println("HOME:", time.Since(start))
}
