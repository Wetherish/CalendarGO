package Calendar

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/pocketbase/pocketbase/models"
)

var db *sql.DB

func StartDB() {
	cfg := mysql.Config{
		User:                 GetDbUser(),
		Passwd:               GetDbPasswd(),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               GetDbName(),
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

func StudentByNameFromDB(name string) ([]Student, error) {
	// An albums slice to hold data from returned rows.
	var Studentlist []Student

	rows, err := db.Query("SELECT * FROM student WHERE student_name = ?", name)
	if err != nil {
		return nil, fmt.Errorf("StudentByStudentName %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var Student Student
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

func GetStudentByIDFromDB(id string) (Student, error) {
	// An album to hold data from the returned row.
	var student Student

	row := db.QueryRow("SELECT * FROM student WHERE id = ?", id)
	if err := row.Scan(&student.Id, &student.Name); err != nil {
		if err == sql.ErrNoRows {
			return student, fmt.Errorf("albumsById %v: no such album", id)
		}
		return student, fmt.Errorf("albumsById %v: %v", id, err)
	}
	return student, nil
}

func DeleteByIDFromDB(id string) (Student, error) {
	student, err := GetStudentByIDFromDB(id)

	_, error := db.Exec("DELETE FROM student WHERE id = ?;", id)

	if error != nil {
		fmt.Println("Error deleteing row: " + err.Error())
		return student, err
	}
	return student, error
}

func DeleteEmployeeFromDB(db *sql.DB, name string) (int64, error) {
	tsql := fmt.Sprintf("DELETE FROM TestSchema.Employees WHERE Name='%s';", name)
	result, err := db.Exec(tsql)
	if err != nil {
		fmt.Println("Error deleting row: " + err.Error())
		return -1, err
	}
	return result.RowsAffected()
}

func AddStudentFromDB(student Student) (int64, error) {
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

func GetAllStudentsFromDB() []Student{

	var StudentList = []Student{}
	
	err := app.Dao().DB().
		Select("id", "Name","Street").
		From("Student").
		All(&StudentList)

	if err != nil {
		fmt.Printf("addAlbum: %v", err)
		return StudentList
	}
		
	return StudentList
}

func AddStudent(student Student)error {
colletion , err := app.Dao().FindCollectionByNameOrId("Student")
if err != nil {
	return err
}
record:= models.NewRecord(colletion)
record.Set("Id",student.Id)
record.Set("Name",student.Name)
record.Set("Street",student.Street)

err = app.Dao().SaveRecord(record)

if err != nil{
	return err
}

return nil
}
