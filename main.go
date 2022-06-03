package main

import (
	"database/sql"
	"fmt"
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
	clmNames := [5]string{"Id", "Path", "Release_date", "Title", "Name"}
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

	fmt.Println(song)

	insertTable(song, song_id)

	http.HandleFunc("/", outputTable)

	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:80", nil)
}
