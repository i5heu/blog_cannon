package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/russross/blackfriday"
)

func PageCacheLoader() {
	start := time.Now()

	ids, err := db.Query("SELECT title,tags,category,text FROM `item` ORDER by timecreate DESC LIMIT 10")
	checkErr(err)

	for ids.Next() {
		var title string
		var tags string
		var category string
		var text string
		_ = ids.Scan(&title, &tags, &category, &text)
		checkErr(err)

		var MainCacheTMP template.HTML

		MainCacheTMP = template.HTML("<article><h1>") + template.HTML(title) + template.HTML("</h1><div class='category'>") + template.HTML(category) + template.HTML("</div><div class='tags'>") + template.HTML(tags) + template.HTML("</div>") + template.HTML(blackfriday.MarkdownCommon([]byte(text))) + template.HTML("</article>")

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

	list := lista{}

	if TMPCACHEWRITE == false {
		list = lista{template.HTML(conf.BlogName), time.Since(start), TMPCACHE[slug]}
	} else if TMPCACHECACHEWRITE == false {
		list = lista{template.HTML(conf.BlogName), time.Since(start), TMPCACHECACHE[slug]}
	} else {
		list = lista{template.HTML(conf.BlogName), time.Since(start), template.HTML("<b> CACHE ERROR <br> Please reload this page</b>")}
	}

	templatesView.Execute(w, list)

	fmt.Println("PAGE:", time.Since(start))
}

func CategoryHandler(w http.ResponseWriter, r *http.Request) {

}
