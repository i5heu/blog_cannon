package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func ArchiveHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	lists := lista{}

	if TMPCACHEWRITE == false {
		lists = lista{template.HTML(conf.BlogName), time.Since(start), TMPCACHE["archive"], false, conf.CoinHiveToken, conf.TrackTrackGo}
	} else if TMPCACHECACHEWRITE == false {
		lists = lista{template.HTML(conf.BlogName), time.Since(start), TMPCACHECACHE["archive"], false, conf.CoinHiveToken, conf.TrackTrackGo}
	} else {
		lists = lista{template.HTML(conf.BlogName), time.Since(start), template.HTML("<b>Please reload this page</b>"), false, conf.CoinHiveToken, conf.TrackTrackGo}
	}

	templatesDesktop.Execute(w, lists)

	fmt.Println("ARCHIVE:", time.Since(start))
}

func ArchiveCacheFunc(foo string) {
	start := time.Now()
	var MainCacheTMP template.HTML

	MainCacheTMP += template.HTML("<style>body{max-width:800px;}</style><table><tr><th>Name</th><th>Category</th><th>Method</th><th>Tags</th><th>Created</th></tr>")

	ids, err := db.Query("SELECT method,title,tags,category,timecreate FROM `item` ORDER by timecreate DESC")
	checkErr(err)

	for ids.Next() {
		var method string
		var title string
		var tags string
		var category string
		var timecreate string
		_ = ids.Scan(&method, &title, &tags, &category, &timecreate)
		checkErr(err)

		slug := category + "/" + title

		MainCacheTMP += template.HTML("<tr><td><a href='/p/") + template.HTML(slug) + template.HTML("'>") + template.HTML(title) + template.HTML("</a></td><td>") + template.HTML(category) + template.HTML("</td><td>") + template.HTML(method) + template.HTML("</td><td>") + template.HTML(tags) + template.HTML("</td><td>") + template.HTML(timecreate) + template.HTML("</td></tr>")
	}

	MainCacheTMP += template.HTML("</table>")

	TMPCACHE[foo] = MainCacheTMP
	fmt.Println("ArchiveCacheFunc:", time.Since(start))
	return
}
