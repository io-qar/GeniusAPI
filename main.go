package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"net/http"
	"html/template"

	_ "github.com/lib/pq"
	"github.com/tidwall/gjson"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "GeniusAPI"
)

var datab *sql.DB

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func dbCon() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	return db
}

func createTable(db *sql.DB, tblName string, clmNames [5]string) {
	_, err := db.Exec(fmt.Sprintf("create table if not exists %s ()", tblName))
	CheckError(err)
	for _, clmName := range clmNames {
		_, err := db.Exec(fmt.Sprintf("alter table %s add column if not exists %s TEXT", tblName, clmName))
		CheckError(err)
	}
}

func insertTable(db *sql.DB, data []gjson.Result, clmNames [5]string, songId int) {
	var i, j int = 0, 0
	var fl bool = false
	var str string = ""

	for ; i < len(clmNames); {
		for ; j < len(data); j++ {
			if (fl) {
				str = "update song_info set " + clmNames[i] + " = '" + data[j].String() + "' where id = '" + strconv.Itoa(songId) + "'"
				_, err := db.Exec(str)
				CheckError(err)
			} else {
				str = "insert into song_info (%s) values ('%s')"
				fl = true
				_, err := db.Exec(fmt.Sprintf(str, clmNames[i], data[j].String()))
				CheckError(err)
			}
			i++
		}
	}
}

func outputTable(w http.ResponseWriter, r *http.Request) {
	rows, err := datab.Query("select * from song_info")
	CheckError(err)
	defer rows.Close()

	songs := []Song{}
	for rows.Next() {
		s := Song{}
		err := rows.Scan(&s.Id, &s.Path, &s.Release_date, &s.Title, &s.Name)
		CheckError(err)
		songs = append(songs, s)
	}

	for _, s := range songs {
		fmt.Println(s.Id, s.Path, s.Release_date, s.Title, s.Name)
	}

	tmpl, err := template.ParseFiles("static/table.html")
	CheckError(err)
	err2 := tmpl.Execute(w, songs)
	CheckError(err2)
}

func main() {
	clmNames := [5]string{"id", "path", "release_date", "title", "name"}
	datab = dbCon()
	defer datab.Close()

	createTable(datab, "song_info", clmNames)
	req, song_id := request()

	results := gjson.GetMany(req, "response.song.id", "response.song.path", "response.song.release_date", "response.song.title", "response.song.album.name")

	insertTable(datab, results, clmNames, song_id)

	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	http.HandleFunc("/table", outputTable)

	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:80", nil)
	
}
