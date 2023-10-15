package Controller

import (
	"Calendar/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTeachers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Models.Teacher{Id: 1, Name: "babcia", Subject: "Math"})
}
