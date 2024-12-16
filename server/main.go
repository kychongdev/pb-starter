package main

import (
	"log"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	// serves static files from the provided public dir (if exists)
	// app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
	// 	e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
	// 	return nil
	// })
	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// serves static files from the provided public dir (if exists)
		se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))

		return se.Next()
	})
	// test

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
