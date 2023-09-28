package main

import (
	"Calendar/Models"
	server "Calendar/Server"
	"calendar/env"
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User:                 env.GetDbName(),
		Passwd:               env.GetDbPasswd(),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               env.GetDbName(),
		AllowNativePasswords: true,
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingerr := db.Ping()
	if pingerr != nil {
		log.Fatal(pingerr)
	}
	fmt.Println("Connected")

	albums, err := albumsByArtist("Maciek")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	alb, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)

	albID, err := addAlbum(Models.Student{
		Name: "Kuba",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)
	server.StartServer()
}

func albumsByArtist(name string) ([]Models.Student, error) {
	// An albums slice to hold data from returned rows.
	var Studentlist []Models.Student

	rows, err := db.Query("SELECT * FROM student WHERE student_name = ?", name)
	if err != nil {
		return nil, fmt.Errorf("StudentByStudentName %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var Student Models.Student
		if err := rows.Scan(&Student.Id, &Student.Name); err != nil {
			return nil, fmt.Errorf("StudentByStudentName %q: %v", name, err)
		}
		Studentlist = append(Studentlist, Student)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("StudentByStudentName %q: %v", name, err)
	}
	return Studentlist, nil
}

func albumByID(id int64) (Models.Student, error) {
	// An album to hold data from the returned row.
	var student Models.Student

	row := db.QueryRow("SELECT * FROM student WHERE id = ?", id)
	if err := row.Scan(&student.Id, &student.Name); err != nil {
		if err == sql.ErrNoRows {
			return student, fmt.Errorf("albumsById %d: no such album", id)
		}
		return student, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return student, nil
}

func addAlbum(student Models.Student) (int64, error) {
	result, err := db.Exec("INSERT INTO student (student_name) VALUES (?)", student.Name)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
