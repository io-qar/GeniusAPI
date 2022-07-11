package main

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/tidwall/gjson"
)

var datab *sql.DB

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	clmNames := [5]string{"id", "path", "release_date", "title", "name"}
	datab = dbCon()
	defer datab.Close()

	createTable("song_info", clmNames)

	req, song_id := request()
	results := gjson.GetMany(req, "response.song.id", "response.song.path", "response.song.release_date", "response.song.title", "response.song.album.name")

	var song = map[string]string {
		clmNames[0]: results[0].String(),
		clmNames[1]: results[1].String(),
		clmNames[2]: results[2].String(),
		clmNames[3]: results[3].String(),
		clmNames[4]: results[4].String(),
	}

	insertTable(song, song_id)

	http.HandleFunc("/", outputTableAll)
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		reqS := r.FormValue("req")
		reqTbl := r.FormValue("category")
		
		tmpl, searchResults := searchTable(reqS, reqTbl)
		
		err2 := tmpl.Execute(w, searchResults)
		CheckError(err2)
	})

	println("Server is listening...");
	http.ListenAndServe("localhost:80", nil)
}