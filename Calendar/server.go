package Calendar

import (
	"net/http"

	"github.com/labstack/echo/v5"
)
func Router(){
	e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
	e.GET("/Students/", GetAllStudent)
	e.GET("/Students/:id", GetStudentByID)
	e.POST("/Students/", PostStudent)
	e.DELETE("/Students/:id", DeleteStudentByID)
	e.Logger.Error(e.Start(":1323"))
}


