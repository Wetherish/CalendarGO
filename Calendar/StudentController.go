package Calendar

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
)

func GetAllStudent(c echo.Context) error {
	// c.JSON(http.StatusOK, GetAllStudentsFromDB())
	c.JSON(http.StatusOK, GetAllStudentsFromDB())
	return nil
}

func PostStudent(c echo.Context) error {
	var newStudent Student
	if err := c.Bind(&newStudent); err != nil {
		return err
	}
	AddStudent(newStudent)
	c.JSON(http.StatusCreated, newStudent)
	return nil
}

func GetStudentByID(c echo.Context) error {

	id := c.PathParam("id")
	student := GetStudentByIDFromDB(id)
	c.JSON(http.StatusOK, student)
	return nil
	//todo error handling
}

func DeleteStudent(c echo.Context) error {
	err := DeleteStudentByIDFromDB(c.PathParam("id"))
	if err != nil {
		fmt.Printf("Delete Student: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return err
	}
	c.JSON(http.StatusOK, "Student deleted")
	return nil
}

func UpdateStudent(c echo.Context) error {
	var newStudent Student
	if err := c.Bind(&newStudent); err != nil {
		return err
	}
	err := UpdateStudentByIDFromDB(c.PathParam("id"), newStudent)
	if err != nil {
		fmt.Printf("Update Student: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return err
	}
	c.JSON(http.StatusOK, "Student updated")
	return nil
}