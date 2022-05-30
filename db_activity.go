package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"github.com/tidwall/gjson"
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

func createTable(tblName string, clmNames [5]string) {
	_, err := datab.Exec(fmt.Sprintf("create table if not exists %s ()", tblName))
	CheckError(err)
	for _, clmName := range clmNames {
		_, err := datab.Exec(fmt.Sprintf("alter table %s add column if not exists %s TEXT", tblName, clmName))
		CheckError(err)
	}
}

func insertTable(data []gjson.Result, clmNames [5]string, songId int) {
	var i, j int = 0, 0
	var fl bool = false
	var str string = ""

	for ; i < len(clmNames); {
		for ; j < len(data); j++ {
			if (fl) {
				str = "update song_info set " + clmNames[i] + " = '" + data[j].String() + "' where id = '" + strconv.Itoa(songId) + "'"
				_, err := datab.Exec(str)
				CheckError(err)
			} else {
				str = "insert into song_info (%s) values ('%s')"
				fl = true
				_, err := datab.Exec(fmt.Sprintf(str, clmNames[i], data[j].String()))
				CheckError(err)
			}
			i++
		}
	}
}