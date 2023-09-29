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
	//todo
}

var StudentList = []Models.Student{
	{Id: 1, Name: "John Coltrane"},
	{Id: 2, Name: "b b"},
}

// func Temp() {
// 	albums, err := dataaccess.StudentByName("maciek")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("Albums found: %v\n", albums)

// 	student, err := dataaccess.GetStudentByID(2)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("Album found: %v\n", student)

// 	studentID, err := dataaccess.AddStudent(Models.Student{
// 		Name: "Kuba",
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("ID of added album: %v\n", studentID)
// }
