package Calendar

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
)

func GetAllLesson(c echo.Context) error {
	c.JSON(http.StatusOK, GetAllLessonFromDB())
	return nil
}

func PostLesson(c echo.Context) error {
	var newLesson Lesson
	if err := c.Bind(&newLesson); err != nil {
		return err
	}
	AddLesson(newLesson)
	c.JSON(http.StatusCreated, newLesson)
	return nil
}

func GetLessonByID(c echo.Context) error {

	id := c.PathParam("id")
	student := GetLessonByIDFromDB(id)
	c.JSON(http.StatusOK, student)
	return nil
	//todo error handling
}

func DeleteLesson(c echo.Context) error {
	err := DeleteLessonByIDFromDB(c.PathParam("id"))
	if err != nil {
		fmt.Printf("Lesson Teacher: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return err
	}
	c.JSON(http.StatusOK, "Lesson deleted")
	return nil
}

func UpdateLesson(c echo.Context) error {
	var newLesson Lesson
	if err := c.Bind(&newLesson); err != nil {
		return err
	}
	err := UpdateLessonByIDFromDB(c.PathParam("id"), newLesson)
	if err != nil {
		fmt.Printf("Update Lesson: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return err
	}
	c.JSON(http.StatusOK, "Lesson updated")
	return nil
}

func GetLessonByStudentID(c echo.Context) error {
	studentID := c.PathParam("id")
	lesson := GetLessonByStudentIDFromDB(studentID)
	c.JSON(http.StatusOK, lesson)
	return nil
}

func GetLessonByTeacherID(c echo.Context) error {
	teacherID := c.PathParam("id")
	lesson := GetLessonByTeacherIDFromDB(teacherID)
	c.JSON(http.StatusOK, lesson)
	return nil
}