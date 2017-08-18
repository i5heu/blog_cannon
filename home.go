package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/russross/blackfriday"
)

type lista struct {
	Rendertime time.Duration
	Articles   template.HTML
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	lists := lista{}

	if TMPCACHEWRITE == false {
		lists = lista{time.Since(start), TMPCACHE["maincache"]}
	} else if TMPCACHECACHEWRITE == false {
		lists = lista{time.Since(start), TMPCACHECACHE["maincache"]}
	} else {
		lists = lista{time.Since(start), template.HTML("<b>Please reload this page</b>")}
	}

	templatesDesktop.Execute(w, lists)

	fmt.Println("HOME:", time.Since(start))
}

func MainCacheFunc(foo string) {
	start := time.Now()
	var MainCacheTMP template.HTML

	ids, err := db.Query("SELECT title,tags,category,text FROM `article` ORDER by timecreate DESC LIMIT 10")
	checkErr(err)

	for ids.Next() {
		var title string
		var tags string
		var category string
		var text string
		_ = ids.Scan(&title, &tags, &category, &text)
		checkErr(err)

		slug := category + "/" + title

		MainCacheTMP += template.HTML("<article><a href='/p/") + template.HTML(slug) + template.HTML("'><h1>") + template.HTML(title) + template.HTML("</h1><a><div class='category'>") + template.HTML(category) + template.HTML("</div><div class='tags'>") + template.HTML(tags) + template.HTML("</div>") + template.HTML(blackfriday.MarkdownCommon([]byte(text))) + template.HTML("</article>")
	}

	TMPCACHE[foo] = MainCacheTMP
	fmt.Println("MainCacheFunc:", time.Since(start))
	return
}
