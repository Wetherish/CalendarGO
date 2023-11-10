package Calendar

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)
var app = pocketbase.New()
func PocketbaseInit() {
    // app := pocketbase.New()

    // serves static files from the provided public dir (if exists)
    app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
        // e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
        e.Router.GET("/Students/", GetAllStudent)
        e.Router.GET("/Students/:id", GetStudentByID)
        e.Router.POST("/Students/", PostStudent)
        e.Router.DELETE("/Students/:id", DeleteStudentByID)
        return nil

    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}
