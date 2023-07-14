package work_offers

import (
	"fmt"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gofrs/uuid"
	"net/http"
	"project_storm/models"
	"project_storm/public"
	"project_storm/templates"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

type WorkOfferHandlers struct {
	tx *pop.Connection
}

var r = render.New(render.Options{
	HTMLLayout:  "application.plush.html",
	TemplatesFS: templates.FS(),
	AssetsFS:    public.FS(),
})

func NewWorkOfferHandlers(tx *pop.Connection) WorkOfferHandlers {
	return WorkOfferHandlers{tx}
}

func (w WorkOfferHandlers) GetWorkOffers(ctx buffalo.Context) error {
	workOffers := models.WorkOffers{}
	err := w.tx.All(&workOffers)
	if err != nil {
		return ctx.Error(500, err)
	}

	ctx.Set("workOffers", workOffers)

	return ctx.Render(http.StatusOK, r.HTML("work_offers/index.plush.html"))
}

func (w WorkOfferHandlers) NewWorkOffer(ctx buffalo.Context) error {
	workOffer := models.WorkOffer{}
	// Recuerda cambiar esto por el ID de la compañía creada
	workOffer.CompanyID = uuid.Must(uuid.NewV4())

	ctx.Set("workOffer", workOffer)

	return ctx.Render(http.StatusOK, r.HTML("work_offers/new.plush.html"))
}

func (w WorkOfferHandlers) CreateWorkOffer(ctx buffalo.Context) error {
	workOffer := models.WorkOffer{}
	if err := ctx.Bind(&workOffer); err != nil {
		return ctx.Error(500, err)
	}

	fmt.Println("workOffer: ", workOffer)

	verrs, err := w.tx.ValidateAndCreate(&workOffer)
	if err != nil {
		return ctx.Error(500, err)
	}

	if verrs.HasAny() {
		ctx.Set("workOffer", workOffer)
		ctx.Set("errors", verrs.Errors)
		return ctx.Render(http.StatusOK, r.HTML("work_offers/new.plush.html"))
	}

	return ctx.Redirect(http.StatusSeeOther, "/admin/work_offers")
}

func (w WorkOfferHandlers) EditWorkOffer(ctx buffalo.Context) error {
	workOffer := models.WorkOffer{}
	if err := w.tx.Find(&workOffer, ctx.Param("id")); err != nil {
		return ctx.Error(500, err)
	}

	ctx.Set("workOffer", workOffer)

	return ctx.Render(http.StatusOK, r.HTML("work_offers/edit.plush.html"))
}

func (w WorkOfferHandlers) UpdateWorkOffer(ctx buffalo.Context) error {
	workOffer := models.WorkOffer{}
	if err := ctx.Bind(&workOffer); err != nil {
		return ctx.Error(500, err)
	}

	fmt.Println("workOffer: ", workOffer.Name)

	err := w.tx.Update(&workOffer)
	if err != nil {
		return ctx.Error(500, err)
	}

	return ctx.Redirect(http.StatusFound, "/admin/work_offers")
}

func (w WorkOfferHandlers) DeleteWorkOffer(ctx buffalo.Context) error {
	workOffer := models.WorkOffer{}
	err := w.tx.Find(&workOffer, ctx.Param("id"))
	if err != nil {
		return ctx.Error(500, err)
	}

	err = w.tx.Destroy(&workOffer)
	if err != nil {
		return ctx.Error(500, err)
	}

	return ctx.Redirect(http.StatusSeeOther, "/admin/work_offers")
}
