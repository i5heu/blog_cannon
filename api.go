package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ApiSTRUCT struct {
	PWD      string
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
		fmt.Fprintf(w, "ERROR")
		checkErr(err)
		return
	}

	if personalpwd == Sjson.PWD {
		fmt.Fprintf(w, "Not Loged in")
	}

	db.Exec("INSERT INTO article(title,category,tags,text) VALUES(?,?,?,?)", ReplaceSpecialChars(Sjson.Title), Sjson.Category, Sjson.Tags, Sjson.Text)

	fmt.Fprintf(w, `{"Status":"OK"}`)
	fmt.Println("ApiHandler:", time.Since(start))
}
