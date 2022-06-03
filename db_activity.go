package main

import (
	"database/sql"
	"fmt"
	// "reflect"
	"strconv"

	// "github.com/tidwall/gjson"
	// "google.golang.org/api/keep/v1"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "GeniusAPI"
)

func dbCon() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	return db
}

func createDb(dbname string) {
	_, err := datab.Exec(fmt.Sprintf("create database '%s'", dbname))
	CheckError(err)
}

func createTable(tblName string, clmNames [5]string) {
	_, err := datab.Exec(fmt.Sprintf("create table if not exists %s ()", tblName))
	CheckError(err)
	for _, clmName := range clmNames {
		_, err := datab.Exec(fmt.Sprintf("alter table %s add column if not exists %s TEXT", tblName, clmName))
		CheckError(err)
	}
}

func insertTable(s map[string]string, songId int) {
	var (
		fl bool = false
		str string = ""
	)

	for key, val := range s {
		if (fl) {
			str = "update song_info set " + key + " = '" + val + "' where Id = '" + strconv.Itoa(songId) + "'"
			_, err := datab.Exec(str)
			CheckError(err)
		} else {
			str = "insert into song_info (%s) values ('%s')"
			fl = true
			_, err := datab.Exec(fmt.Sprintf(str, key, val))
			CheckError(err)
		}
	}
}