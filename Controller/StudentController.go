package Controller

import (
	"Calendar/Models"
	dataaccess "Calendar/data-access"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, dataaccess.GetAllStudents())
}

func PostAlbums(c *gin.Context) {
	var newStudent Models.Student
	if err := c.BindJSON(&newStudent); err != nil {
		return
	}
	dataaccess.AddStudent(newStudent)
	c.IndentedJSON(http.StatusCreated, newStudent)
}

func GetAlbumsByID(c *gin.Context) {

	id := c.Param("id")
	student, err := dataaccess.GetStudentByID(id)
	if err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, student)
}

func DeleteByID(c *gin.Context) {
	// //todo
	id := c.Param("id")
	student, err := dataaccess.DeleteByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNoContent, "no such data")
	}
	c.IndentedJSON(http.StatusOK, student)
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
