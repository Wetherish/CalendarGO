package Calendar

import (
	"fmt"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/models"
)

func GetAllTeachersFromDB() []Teacher {
	teachersList := []Teacher{}
	err := app.Dao().DB().
		Select("id", "Name", "Street").
		From("Teacher").
		All(&teachersList)

	if err != nil {
		fmt.Printf("addAlbum: %v", err)
		return teachersList
	}
	return teachersList
}

func AddTeacher(teacher Teacher) error {
	colletion, err := app.Dao().FindCollectionByNameOrId("Teacher")
	if err != nil {
		return err
	}
	record := models.NewRecord(colletion)
	record.Set("Id", teacher.Id)
	record.Set("Name", teacher.Name)
	record.Set("Street", teacher.Street)

	err = app.Dao().SaveRecord(record)

	if err != nil {
		return err
	}
	return nil
}

func GetTeacherByIDFromDB(id string) (Teacher, error) {
	teacher := Teacher{}
	err := app.Dao().DB().
		Select("id", "Name", "Street").
		From("Teacher").
		Where(dbx.NewExp("id = {:id}", dbx.Params{"id": id})).
		One(&teacher)

	if err != nil {
		fmt.Printf("addAlbum: %v", err)
		return teacher, err
	}
	return teacher, nil
}

func DeleteTeacherByIDFromDB(id string) error {
	record, err := app.Dao().FindRecordById("Teacher", id)
	if err != nil {
		return err
	}

	if err := app.Dao().DeleteRecord(record); err != nil {
		return err
	}
	return nil
}

func UpdateTeacherByIDFromDB(id string, teacher Teacher) error {
	record, err := app.Dao().FindRecordById("Teacher", id)
	if err != nil {
		return err
	}
	record.Set("Name", teacher.Name)
	record.Set("Street", teacher.Street)

	if err := app.Dao().SaveRecord(record); err != nil {
		return err
	}
	return nil
}