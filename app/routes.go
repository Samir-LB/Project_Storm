package app

import (
	"net/http"

	"Project_Storm/app/actions/home"
	"Project_Storm/app/middleware"
	"Project_Storm/public"

	"github.com/gobuffalo/buffalo"
)

// SetRoutes for the application
func setRoutes(root *buffalo.App) {
	root.Use(middleware.RequestID)
	root.Use(middleware.Database)
	root.Use(middleware.ParameterLogger)
	root.Use(middleware.CSRF)

	root.GET("/", home.Index)
	root.POST("/send-crendetials", home.VerifyUser).Name("sendCrendentials")
	root.ServeFiles("/", http.FS(public.FS()))
}
