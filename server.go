package main

import (
	// "fmt"
	"net/http"
	"html/template"
)

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

	// for _, s := range songs {
	// 	fmt.Println(s.Id, s.Path, s.Release_date, s.Title, s.Name)
	// }

	tmpl, err := template.ParseFiles("static/table.html")
	CheckError(err)
	err2 := tmpl.Execute(w, songs)
	CheckError(err2)
}