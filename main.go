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

	insertTable(results, clmNames, song_id)

	http.HandleFunc("/", outputTable)

	// fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:80", nil)
}
