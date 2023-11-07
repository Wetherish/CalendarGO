package Controller

import (
	"Calendar/Models"
	dataaccess "Calendar/data-access"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllStudent(c echo.Context) error{
	c.JSON(http.StatusOK, dataaccess.GetAllStudents())
	return nil
}

func PostStudent(c echo.Context) error{
	var newStudent Models.Student
	if err := c.Bind(&newStudent); err != nil {
		return err
	}
	dataaccess.AddStudent(newStudent)
	c.JSON(http.StatusCreated, newStudent)
	return nil
}

func GetStudentByID(c echo.Context) error{

	id := c.Param("id")
	student, err := dataaccess.GetStudentByID(id)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, student)
	return nil
}

func DeleteStudentByID(c echo.Context) error{
	// //todo
	id := c.Param("id")
	student, err := dataaccess.DeleteByID(id)
	if err != nil {
		c.JSON(http.StatusNoContent, "no such data")
		return err
	}
	c.JSON(http.StatusOK, student)
	return nil
}

// func deleteMovie(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range movies {
// 		if item.ID == params["id"] {
// 			movies = append(movies[:index], movies[index+1:]...)
// 			break
// 		}
// 	}
// }

func GetStudentNamebyId(id int) string {
	//todo
	return "name"
}
