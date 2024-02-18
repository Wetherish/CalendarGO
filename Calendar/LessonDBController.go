package Calendar

import (
	"fmt"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/models"
)

func FindAllLesson() []Lesson {
	lessonList := []Lesson{}
	err := app.Dao().DB().
		Select("id", "Subject", "Date", "Place", "StudentID", "TeacherID", "DataStart", "DataEnd").
		From("Lesson").
		All(&lessonList)

	if err != nil {
		fmt.Printf("addAlbum: %v", err)
		return lessonList
	}
	return lessonList
}

func AddLesson(Lesson Lesson) error {

	collection, err := app.Dao().FindCollectionByNameOrId("Lesson")
	if err != nil {
		return err
	}
	record := models.NewRecord(collection)
	record.Set("Id", Lesson.Id)
	record.Set("Subject", Lesson.Subject)
	record.Set("TeacherId", Lesson.TeacherId)
	record.Set("StudentId", Lesson.StudentId)
	record.Set("Place", Lesson.Place)
	record.Set("Date", Lesson.Date)
	err = app.Dao().SaveRecord(record)

	if err != nil {
		return err
	}

	return nil
}

func FindLessonByID(id string) Teacher {
	teacher := Teacher{}
	err := app.Dao().DB().
		Select("StudentId", "Place", "Date", "TeacherId", "Subject", "Id").
		From("Lesson").
		Where(dbx.NewExp("id = {:id}", dbx.Params{"id": id})).
		One(&teacher)

	if err != nil {
		fmt.Printf("addAlbum: %v", err)
		return teacher
	}
	return teacher
}

func DeleteLessonByID(id string) error {
	record, err := app.Dao().FindRecordById("Teacher", id)
	if err != nil {
		return err
	}

	if err := app.Dao().DeleteRecord(record); err != nil {
		return err
	}
	return nil
}

func UpdateLessonByID(id string, newLesson Lesson) error {
	record, err := app.Dao().FindRecordById("Lesson", id)
	if err != nil {
		return err
	}
	record.Set("Id", newLesson.Id)
	record.Set("Subject", newLesson.Subject)
	record.Set("TeacherId", newLesson.TeacherId)
	record.Set("StudentId", newLesson.StudentId)
	record.Set("Place", newLesson.Place)
	record.Set("Date", newLesson.Date)
	if err := app.Dao().SaveRecord(record); err != nil {
		return err
	}
	return nil
}

func FindLessonByStudentID(id string) []Lesson {
	lessonList := []Lesson{}
	err := app.Dao().DB().
		Select("id", "Subject", "Date", "Place", "StudentID", "TeacherID").
		From("Lesson").
		Where(dbx.NewExp("StudentID = {:id}", dbx.Params{"id": id})).
		All(&lessonList)

	if err != nil {
		fmt.Printf("addAlbum: %v", err)
		return lessonList
	}
	return lessonList
}

func FindLessonByTeacherID(id string) []Lesson {
	lessonList := []Lesson{}
	err := app.Dao().DB().
		Select("id", "Subject", "Date", "Place", "StudentID", "TeacherID").
		From("Lesson").
		Where(dbx.NewExp("TeacherID = {:id}", dbx.Params{"id": id})).
		All(&lessonList)

	if err != nil {
		fmt.Printf("addAlbum: %v", err)
		return lessonList
	}
	return lessonList
}
