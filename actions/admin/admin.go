package admin

import (
	"net/http"
	"project_storm/public"
	"project_storm/templates"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
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

func Home(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("dashboard/admin_dashboard.plush.html"))
}
