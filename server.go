package main

import (
	"net/http"
	"html/template"
)

func outputTableAll(w http.ResponseWriter, r *http.Request) {
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

	tmpl, err := template.ParseFiles("static/table.html")
	CheckError(err)
	err2 := tmpl.Execute(w, songs)
	CheckError(err2)
}