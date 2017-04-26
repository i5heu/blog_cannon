package main

import (
	"database/sql"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Global sql.DB to access the database by all handlers
var db *sql.DB
var err error

func main() {
	// Create an sql.DB and check for errors
	db, err = sql.Open("mysql", "USER:PASSWORD@/blog_cannon")
	if err != nil {
		panic(err.Error())
	}
	// sql.DB should be long lived "defer" closes it once this function ends
	defer db.Close()

	db.Exec("CREATE TABLE IF NOT EXISTS article (id INT NOT NULL AUTO_INCREMENT, title VARCHAR (128) NOT NULL default 'NO TITLE', text longtext, PRIMARY KEY (id),FULLTEXT (title,text))")

	http.HandleFunc("/index/", IndexHandler)
	http.HandleFunc("/newentry", NewentryHandler)
	http.HandleFunc("/p/", ViewHandler)
	http.HandleFunc("/s/", SearchHandler)
	http.HandleFunc("/", IndexHandler2)
	http.ListenAndServe(":8080", nil)
}

func IndexHandler2(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/index", 302)

}

func readfiles() (tmp []string) {
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		tmp = append(tmp, strings.TrimSuffix(f.Name(), filepath.Ext(f.Name())))
	}

	return tmp
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
