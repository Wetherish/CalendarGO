package Calendar

import (
	"log"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

var app = pocketbase.New()

func PocketbaseInit() {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), true))
		e.Router.GET("/api", GetApisList)

		e.Router.GET("/Students/", GetAllStudent, apis.ActivityLogger(app) )
		e.Router.GET("/Teachers/", GetAllTeachers, apis.ActivityLogger(app))
		e.Router.GET("/Lessons/", GetAllLesson, apis.ActivityLogger(app))

		e.Router.POST("/Students/", PostStudent, apis.ActivityLogger(app))
		e.Router.POST("/Teachers/", PostTeacher, apis.ActivityLogger(app))
		e.Router.POST("/Lessons/", PostLesson, apis.ActivityLogger(app))

		e.Router.DELETE("/Students/:id", DeleteStudent, apis.ActivityLogger(app))
		e.Router.DELETE("/Teachers/:id", DeleteTeacher, apis.ActivityLogger(app))
		e.Router.DELETE("/Lesson/:id", DeleteLesson, apis.ActivityLogger(app))

		e.Router.GET("/Students/:id", GetStudentByID, apis.ActivityLogger(app))
		e.Router.GET("/Teachers/:id", GetTeacherByID, apis.ActivityLogger(app))
		e.Router.GET("/Lesson/:id", GetLessonByID, apis.ActivityLogger(app))

		//todo e.Router.PATCH("/Students/:id", PatchStudentByID)
		//todo e.Router.PATCH("/Teachers/:id", PatchTeacherByID)
		//todo e.Router.PATCH("/Lessons/:id", PatchLessonByID)

		return nil

	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
