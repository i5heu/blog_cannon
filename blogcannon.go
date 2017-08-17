package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var HtmlStructHeader string = templatefolder + `/header.html`
var HtmlStructFooter string = templatefolder + `/footer.html`

var templatesDesktop = template.Must(template.ParseFiles("./template/home.html", HtmlStructHeader, HtmlStructFooter))

var fs = http.FileServer(http.Dir("static"))

var db *sql.DB
var err error
var TMPCACHE = make(map[string]template.HTML)
var TMPCACHECACHE = make(map[string]template.HTML)
var TMPCACHEWRITE bool = false
var TMPCACHECACHEWRITE bool = false

func main() {
	fmt.Println("HALLO")

	db, err = sql.Open("mysql", dblogin)
	db.SetConnMaxLifetime(time.Second * 2)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(25)

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic
	}

	if err != nil {
		panic(err.Error())
	}
	// sql.DB should be long lived "defer" closes it once this function ends
	defer db.Close()

	db.Exec("CREATE TABLE IF NOT EXISTS `article` ( `id` int(10) unsigned NOT NULL AUTO_INCREMENT, `title` varchar(100) NOT NULL DEFAULT 'NO TITLE', `tags` varchar(500) NOT NULL DEFAULT 'NO TAGS',`category` varchar(100) NOT NULL DEFAULT 'main', `timecreate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, `text` longtext NOT NULL, PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1")

	go func() {
		for {
			TMPCACHEWRITE = true
			time.Sleep(500 * time.Millisecond)
			MainCacheFunc("maincache")
			PageCacheLoader()
			ArchiveCacheFunc("archive")
			TMPCACHEWRITE = false
			time.Sleep(500 * time.Millisecond)

			TMPCACHECACHEWRITE = true
			time.Sleep(500 * time.Millisecond)

			for key, value := range TMPCACHE {
				TMPCACHECACHE[key] = value
			}

			TMPCACHECACHEWRITE = false

			time.Sleep(120 * time.Second)
		}
	}()

	http.HandleFunc("/p/", PageHandler)
	http.HandleFunc("/archive", ArchiveHandler)
	http.HandleFunc("/api", ApiHandler)
	http.HandleFunc("/favicon.ico", FaviconHandler)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":8083", nil)
}

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/static/favicon/favicon.ico", 302)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("\033[0;31m", err, "\033[0m")
		err = nil
	}
}
