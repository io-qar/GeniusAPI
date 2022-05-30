package main

import (
	// "database/sql"
	// "fmt"
	// "net/http"
	// "html/template"
)

// func outputTable(w http.ResponseWriter, r *http.Request) {
// 	rows, err := db.Query("select * from song_info")
// 	CheckError(err)

// 	songs := []Song{}
// 	for rows.Next() {
// 		s := Song{}
// 		err := rows.Scan(&s.artist_names, &s.title, &s.song_id, &s.path, &s.album_title)
			
// 		if err != nil{
// 			fmt.Println(err)
// 			continue
// 		}	
// 		songs = append(songs, s)
// 	}
// 	rows.Close()

// 	for _, s := range songs {
// 		fmt.Println(s.artist_names, s.title, s.song_id, s.path, s.album_title)
// 	}

// 	tmpl, _ := template.ParseFiles("static/table.html")
// 	tmpl.Execute(w, songs)

// 		// fmt.Fprint(w, `
// 		// 	<table>
// 		// 		<tr>
// 		// 			<th>test</th>
// 		// 			<th>test2</th>
// 		// 		</tr>
// 		// 	</table>
// 		// `)
	
// }