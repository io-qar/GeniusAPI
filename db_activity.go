package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"html/template"
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

// func createDb(dbname string) {
// 	_, err := datab.Exec(fmt.Sprintf("create database '%s'", dbname))
// 	CheckError(err)
// }

func createTable(tblName string, clmNames [5]string) {
	_, err := datab.Exec(fmt.Sprintf("create table if not exists %s ()", tblName))
	CheckError(err)
	for _, clmName := range clmNames {
		_, err := datab.Exec(fmt.Sprintf("alter table %s add column if not exists %s TEXT", tblName, clmName))
		CheckError(err)
		_, err = datab.Exec(fmt.Sprintf("alter table %s alter column %s set not null", tblName, clmName))
		CheckError(err)
		_, err = datab.Exec(fmt.Sprintf("alter table %s alter column %s set default '--'", tblName, clmName))
	}
}

func insertTable(s map[string]string, songId int) {
	var (
		fl bool = false
		str string = ""
	)

	for key, val := range s {
		if val == "" {
			continue
		} else {
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
}

func searchTable(reqS string, reqTbl string) (*template.Template, []Song) {
	var str string = "select * from song_info where " + reqTbl + " like '%" + reqS + "%'"
	rows, err := datab.Query(str)
	CheckError(err)
	defer rows.Close()

	searchResults := []Song{}
	for rows.Next() {
		s := Song{}
		err := rows.Scan(&s.Id, &s.Path, &s.Release_date, &s.Title, &s.Name)
		CheckError(err)
		searchResults = append(searchResults, s)
	}

	tmpl, err := template.ParseFiles("static/searchResults.html")
	CheckError(err)

	return tmpl, searchResults
}