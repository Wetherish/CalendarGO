package dataaccess

import (
	"Calendar/Models"
	"Calendar/env"
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func StartDB() {
	cfg := mysql.Config{
		User:                 env.GetDbUser(),
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

}

func StudentByName(name string) ([]Models.Student, error) {
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

func GetStudentByID(id string) (Models.Student, error) {
	// An album to hold data from the returned row.
	var student Models.Student

	row := db.QueryRow("SELECT * FROM student WHERE id = ?", id)
	if err := row.Scan(&student.Id, &student.Name); err != nil {
		if err == sql.ErrNoRows {
			return student, fmt.Errorf("albumsById %v: no such album", id)
		}
		return student, fmt.Errorf("albumsById %v: %v", id, err)
	}
	return student, nil
}

func DeleteByID(id string) (Models.Student, error) {
	student, err := GetStudentByID(id)

	_, error := db.Exec("DELETE FROM student WHERE id = ?;", id)

	if error != nil {
		fmt.Println("Error deleteing row: " + err.Error())
		return student, err
	}
	return student, error
}

func DeleteEmployee(db *sql.DB, name string) (int64, error) {
	tsql := fmt.Sprintf("DELETE FROM TestSchema.Employees WHERE Name='%s';", name)
	result, err := db.Exec(tsql)
	if err != nil {
		fmt.Println("Error deleting row: " + err.Error())
		return -1, err
	}
	return result.RowsAffected()
}

func AddStudent(student Models.Student) (int64, error) {
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

func GetAllStudents() []Models.Student {
	tableName := "student"
	var StudentList = []Models.Student{}
	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			Id   int
			Name string
		)
		err := rows.Scan(&Id, &Name)
		if err != nil {
			log.Fatal(err)
		}
		StudentList = append(StudentList, Models.Student{Id: strconv.Itoa(Id), Name: Name})
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return StudentList
}
