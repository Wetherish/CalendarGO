package Calendar

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
)

func GetAllTeachers(c echo.Context) error {
	c.JSON(http.StatusOK, GetAllTeachersFromDB())
	return nil
}

func PostTeacher(c echo.Context) error {
	var newTeacher Teacher
	if err := c.Bind(&newTeacher); err != nil {
		return err
	}
	AddTeacher(newTeacher)
	c.JSON(http.StatusCreated, newTeacher)
	return nil
}

func GetTeacherByID(c echo.Context) error {
	Student, err := GetTeacherByIDFromDB(c.PathParam("id"))
	if err != nil {
		fmt.Printf("Get Teacher: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return err
	}
	c.JSON(http.StatusOK, Student)
	return nil
}

func DeleteTeacher(c echo.Context) error {
	err := DeleteTeacherByIDFromDB(c.PathParam("id"))
	if err != nil {
		fmt.Printf("Delete Teacher: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return err
	}
	c.JSON(http.StatusOK, "Teacher deleted")
	return nil
}

func UpdateTeacher(c echo.Context) error {
	var newTeacher Teacher
	if err := c.Bind(&newTeacher); err != nil {
		return err
	}
	err := UpdateTeacherByIDFromDB(c.PathParam("id"), newTeacher)
	if err != nil {
		fmt.Printf("Update Teacher: %v", err)
		c.JSON(http.StatusBadRequest, err)
		return err
	}
	c.JSON(http.StatusOK, "Teacher updated")
	return nil
}