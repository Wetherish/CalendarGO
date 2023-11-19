package Calendar

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

var app = pocketbase.New()

func PocketbaseInit() {
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		e.Router.GET("/Students/", GetAllStudent)
		e.Router.POST("/Students/", PostStudent)

		e.Router.GET("/Teachers/", GetAllTeachers)
		e.Router.POST("/Teachers/", PostTeacher)

		e.Router.GET("/Lessons/", GetAllLesson)
		e.Router.POST("/Lessons/", PostLesson)

		e.Router.DELETE("/Students/:id", DeleteStudent) 
		e.Router.DELETE("/Teachers/:id", DeleteTeacher)
		e.Router.DELETE("/Lesson/:id", DeleteLesson)  

		e.Router.GET("/Students/:id", GetStudentByID)
		e.Router.GET("/Teachers/:id", GetTeacherByID)
		e.Router.GET("/Lesson/:id", GetLessonByID)

		//todo e.Router.PATCH("/Students/:id", PatchStudentByID)
		//todo e.Router.PATCH("/Teachers/:id", PatchTeacherByID)
		//todo e.Router.PATCH("/Lessons/:id", PatchLessonByID)

		return nil

	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
