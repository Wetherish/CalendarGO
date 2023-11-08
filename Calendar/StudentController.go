package Calendar

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func GetAllStudent(c echo.Context) error{
	c.JSON(http.StatusOK, GetAllStudentsFromDB())
	return nil
}

func PostStudent(c echo.Context) error{
	var newStudent Student
	if err := c.Bind(&newStudent); err != nil {
		return err
	}
	AddStudentFromDB(newStudent)
	c.JSON(http.StatusCreated, newStudent)
	return nil
}

func GetStudentByID(c echo.Context) error{

	id := c.PathParam("id")
	student, err := GetStudentByIDFromDB(id)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, student)
	return nil
}

func DeleteStudentByID(c echo.Context) error{
	// //todo
	id := c.PathParam("id")
	student, err := DeleteByIDFromDB(id)
	if err != nil {
		c.JSON(http.StatusNoContent, "no such data")
		return err
	}
	c.JSON(http.StatusOK, student)
	return nil
}

func GetStudentNamebyId(id int) string {
	//todo
	return "name"
}
