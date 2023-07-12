package login

import (
	"net/http"
	"project_storm/models"
	"project_storm/public"
	"project_storm/templates"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/pop/v6"
)

var r = render.New(render.Options{
	// HTML layout to be used for all HTML requests:
	HTMLLayout: "application.plush.html",

	// fs.FS containing templates
	TemplatesFS: templates.FS(),

	// fs.FS containing assets
	AssetsFS: public.FS(),

	// Add template helpers here:
	Helpers: render.Helpers{
		// for non-bootstrap form helpers uncomment the lines
		// below and import "github.com/gobuffalo/helpers/forms"
		// forms.FormKey:     forms.Form,
		// forms.FormForKey:  forms.FormFor,
	},
})

func Index(c buffalo.Context) error {
	userName := "Hola"
	c.Set("hola", userName)
	return c.Render(http.StatusOK, r.HTML("login/index.plush.html"))

}

func VerifyUser(c buffalo.Context) error {
	egresado := &models.Egresado{}
	return c.Redirect(http.StatusSeeOther, "/admin/home")
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
	return c.Render(400, r.HTML("login/index.plush.html"))
}
