package Controller

import (
	"Calendar/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, StudentList)
}

func PostAlbums(c *gin.Context) {
	var newStudent Models.Student
	if err := c.BindJSON(&newStudent); err != nil {
		return
	}
	StudentList = append(StudentList, newStudent)
	c.IndentedJSON(http.StatusCreated, newStudent)
}

func GetAlbumsByID(c *gin.Context) {

}

var StudentList = []Models.Student{
	{Id: 1, Name: "John Coltrane", Grade: 5},
	{Id: 2, Name: "b b", Grade: 2},
}
