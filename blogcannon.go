package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/i5heu/gosocial"
)

var HtmlStructHeader string
var HtmlStructFooter string

var fs = http.FileServer(http.Dir("static"))

var db *sql.DB
var err error
var TMPCACHE = make(map[string]template.HTML)
var TMPCACHECACHE = make(map[string]template.HTML)
var TMPCACHEWRITE bool = false
var TMPCACHECACHEWRITE bool = false

var namespaceView, templatesView, templatesDesktop *template.Template

type Config struct {
	Dblogin                  string
	Guestmode                bool
	AdminPWD                 string
	GuestPWD                 string
	Templatefolder           string
	AdminHASH                string
	BlogName                 string
	JabberNotification       bool
	JabberHost               string
	JabberUser               string
	JabberPassword           string
	JabberServerName         string
	JabberJIDreciever        string
	JabberInsecureSkipVerify bool
}

var conf Config // Standart SORCE for Password etc.

func main() {
	fmt.Println("Blog starting....")

	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		fmt.Println(err)
		return
	}

	var foo = sha256.Sum256([]byte(conf.AdminPWD)) //sha256 Parser for Password Token
	conf.AdminHASH = hex.EncodeToString(foo[:])

	HtmlStructHeader = conf.Templatefolder + `/header.html`
	HtmlStructFooter = conf.Templatefolder + `/footer.html`

	namespaceView = template.Must(template.ParseFiles("./template/layout.html"))
	templatesView = template.Must(template.ParseFiles("./template/layout.html"))
	templatesDesktop = template.Must(template.ParseFiles("./template/layout.html"))

	// ################ END CONFIG ###########################

	fmt.Println("Blog START")

	db, err = sql.Open("mysql", conf.Dblogin)
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

	db.Exec("CREATE TABLE IF NOT EXISTS `article` ( `id` int(10) unsigned NOT NULL AUTO_INCREMENT, `title` varchar(100) NOT NULL DEFAULT 'NO TITLE', `tags` varchar(500) NOT NULL DEFAULT 'NO TAGS',`category` varchar(100) NOT NULL DEFAULT 'main', `timecreate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, `text` longtext NOT NULL, PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1")
	gosocial.Init(db)

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
	http.HandleFunc("/gosocial", GoSocial)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":8083", nil)
}

func GoSocial(w http.ResponseWriter, r *http.Request) {
	foo, title, text := gosocial.ApiHandler(w, r, conf.AdminHASH)

	if foo == "WriteComment" {
		bar := "⚐ WriteComment \n###Title###-->\n" + title + "\n\n###Comment###-->\n" + text
		sendXMPP(bar)
	}
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
