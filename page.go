package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type view struct {
	Rendertime time.Duration
	Articles   template.HTML
}

var namespaceView = template.Must(template.ParseFiles("./template/home.html", HtmlStructHeader, HtmlStructFooter))
var templatesView = template.Must(template.ParseFiles("./template/home.html", HtmlStructHeader, HtmlStructFooter))

func PageCacheLoader() {
	start := time.Now()

	ids, err := db.Query("SELECT title,tags,category,text FROM `article` ORDER by timecreate DESC LIMIT 10")
	checkErr(err)

	for ids.Next() {
		var title string
		var tags string
		var category string
		var text string
		_ = ids.Scan(&title, &tags, &category, &text)
		checkErr(err)

		var MainCacheTMP template.HTML

		MainCacheTMP = template.HTML("<article><h1>") + template.HTML(title) + template.HTML("</h1><div class='category'>") + template.HTML(category) + template.HTML("</div><div class='tags'>") + template.HTML(tags) + template.HTML("</div>") + template.HTML(text) + template.HTML("</article>")

		foo := category + "/" + title

		TMPCACHE[foo] = MainCacheTMP
	}

	fmt.Println("PageCacheLoader:", time.Since(start))
	return
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	u, err := url.Parse(r.URL.Path)

	checkErr(err)
	encodetpath1 := strings.Split(u.Path, "/")

	if len(encodetpath1) == 2 {
		CategoryHandler(w, r)
		return
	}
	if len(encodetpath1) < 3 {
		fmt.Fprintf(w, "ERROR 404")
		return
	}

	slug := encodetpath1[2] + "/" + encodetpath1[3]

	views := view{}

	if TMPCACHEWRITE == false {
		views = view{time.Since(start), TMPCACHE[slug]}
	} else if TMPCACHECACHEWRITE == false {
		views = view{time.Since(start), TMPCACHECACHE[slug]}
	} else {
		views = view{time.Since(start), template.HTML("<b> CACHE ERROR <br> Please reload this page</b>")}
	}

	templatesView.Execute(w, views)

	fmt.Println("PAGE:", time.Since(start))
}

func CategoryHandler(w http.ResponseWriter, r *http.Request) {

}
