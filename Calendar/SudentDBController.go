package Calendar

import (
	"fmt"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/models"
)

func FindAllStudents() []Student {

	var StudentList = []Student{}

	err := app.Dao().DB().
		Select("id", "Name", "Street").
		From("Student").
		All(&StudentList)

	if err != nil {
		fmt.Printf("addAlbum: %v", err)
		return StudentList
	}

	return StudentList
}

func AddStudent(student Student) error {
	colletion, err := app.Dao().FindCollectionByNameOrId("Student")
	if err != nil {
		return err
	}
	record := models.NewRecord(colletion)
	record.Set("Id", student.Id)
	record.Set("Name", student.Name)
	record.Set("Street", student.Street)

	err = app.Dao().SaveRecord(record)

	if err != nil {
		return err
	}

	return nil
}

func FindStudentByID(id string) Teacher {
	teacher := Teacher{}
	err := app.Dao().DB().
		Select("id", "Name", "Street").
		From("Student").
		Where(dbx.NewExp("id = {:id}", dbx.Params{"id": id})).
		One(&teacher)

	if err != nil {
		fmt.Printf("addAlbum: %v", err)
		return teacher
	}
	return teacher
}

func DeleteStudentByID(id string) error {
	record, err := app.Dao().FindRecordById("Student", id)
	if err != nil {
		return err
	}

	if err := app.Dao().DeleteRecord(record); err != nil {
		return err
	}
	return nil
}

func UpdateStudentByID(id string, student Student) error {
	record, err := app.Dao().FindRecordById("Student", id)
	if err != nil {
		return err
	}

	record.Set("Name", student.Name)
	record.Set("Street", student.Street)

	if err := app.Dao().SaveRecord(record); err != nil {
		return err
	}
	return nil
}
