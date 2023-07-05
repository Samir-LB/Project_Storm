package home

import (
	"Project_Storm/app/models"
	"Project_Storm/app/render"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
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
	egresado := &models.Egresado{}
	if err := c.Bind(egresado); err != nil {
		return err
	}

	// Realizar validación de existencia del usuario
	tx := c.Value("tx").(*pop.Connection)
	exists, err := tx.Where("correo = ?", egresado.Correo).Where("clave = ?", egresado.Clave).Exists(egresado)
	if err != nil {
		return err
	}
	if exists {
		// El usuario ya existe, redirigir al home
		return c.Redirect(302, "/")
	}

	// mosstrar error
	c.Set("error", "El usuario ya existe.")
	return c.Render(400, r.HTML("home/index.plush.html"))
}
