package Controller

import (
	"Calendar/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLessons(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Models.Lesson{})
}
