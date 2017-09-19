package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ApiSTRUCT struct {
	PWD      string
	Method   string
	Title    string
	Category string
	Tags     string
	Text     string
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	decoder := json.NewDecoder(r.Body)
	var Sjson ApiSTRUCT
	errSearch := decoder.Decode(&Sjson)
	if errSearch != nil {
		fmt.Println(errSearch)
		fmt.Fprintf(w, "ERROR")
		checkErr(err)
		return
	}

	if conf.AdminHASH != Sjson.PWD {
		fmt.Fprintf(w, `{"Status":"Not Loged in"}`)
		fmt.Println("ApiHandler-Not Loged in:", time.Since(start))
		return
	}

	if Sjson.Method == "article" || Sjson.Method == "snippet" || Sjson.Method == "link" {
		db.Exec("INSERT INTO item(method,title,category,tags,text) VALUES(?,?,?,?,?)", Sjson.Method, ReplaceSpecialChars(Sjson.Title), Sjson.Category, Sjson.Tags, Sjson.Text)
	} else {
		fmt.Fprintf(w, `{"Status":"No Valid Method"}`)
		fmt.Println("ApiHandler-No Valid Method:", time.Since(start))
		return
	}

	//if Sjson.Method ==

	fmt.Fprintf(w, `{"Status":"OK"}`)
	fmt.Println("ApiHandler:", time.Since(start))
}
