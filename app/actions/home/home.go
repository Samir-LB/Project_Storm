package home

import (
	"Project_Storm/app/render"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

var (
	// r is a buffalo/render Engine that will be used by actions
	// on this package to render render HTML or any other formats.
	r = render.Engine
)

func Index(c buffalo.Context) error {
	userName := "Hola"
	c.Set("hola", userName)
	return c.Render(http.StatusOK, r.HTML("home/index.plush.html"))
}

func VerifyUser(c buffalo.Context) error {
	userName := c.Value("Username")
	c.Set("hola", userName)
	return c.Redirect(http.StatusOK, "dashboard/dashboard.plush.html")
}
