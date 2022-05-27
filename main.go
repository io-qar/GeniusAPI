package main

import (
	"database/sql"
	"fmt"
	"strconv"

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
		// _, err2 := db.Exec(fmt.Sprintf("alter table %s add primary key (id)", tblName))
		// CheckError(err2)
	}
}

func insertTable(db *sql.DB, data []gjson.Result, clmNames [5]string, songId ...int) {
	var i, j int = 0, 0
	var fl bool = false
	var str = ""

	for ; i < len(clmNames); {
		for ; j < len(data); {
			if (fl) {
				str = "update song_info set " + clmNames[i] + " = '" + data[j].String() + "' where id = '" + strconv.Itoa(songId[0]) + "'"
				_, err := db.Exec(str)
				CheckError(err)
			} else {
				str = "insert into song_info (%s) values ('%s')"
				fl = true
				_, err := db.Exec(fmt.Sprintf(str, clmNames[i], data[j].String()))
				CheckError(err)
			}
			i++
			j++
		}
	}
}

func main() {
	clmNames := [5]string{"id", "path", "release_date", "title", "name"}
	db := dbCon()

	createTable(db, "song_info", clmNames)
	req, song_id := request()

	results := gjson.GetMany(req, "response.song.id", "response.song.path", "response.song.release_date", "response.song.title", "response.song.album.name")

	insertTable(db, results, clmNames, song_id)

	db.Close()
}
