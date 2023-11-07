package server

import (
	"Calendar/Controller"
	"net/http"

	"github.com/labstack/echo/v4"
)
func Router(){
	e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
	e.GET("/Students/", Controller.GetAllStudent)
	e.GET("/Students/:id", Controller.GetStudentByID)
	e.POST("/Students/", Controller.PostStudent)
	e.DELETE("/Students/:id", Controller.DeleteStudentByID)
	e.Logger.Error(e.Start(":1323"))
}


