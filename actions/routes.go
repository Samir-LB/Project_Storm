package actions

import (
	"project_storm/actions/admin"
	"project_storm/actions/login"
	workOffers "project_storm/actions/work_offers"
	"project_storm/models"

	"github.com/gobuffalo/buffalo"
)

func SetRoutes(app *buffalo.App) {
	app.GET("/", login.Index)
	app.POST("/send-crendetials", login.VerifyUser).Name("sendCrendentials")

	adminRoutes := app.Group("/admin")
	adminRoutes.GET("/home", admin.Home)

	// Work offers routes
	tx := models.DB

	workOffersHandlers := workOffers.NewWorkOfferHandlers(tx)
	adminRoutes.GET("/work_offers", workOffersHandlers.GetWorkOffers)
	adminRoutes.GET("/work_offers/new", workOffersHandlers.NewWorkOffer)
	adminRoutes.POST("/work_offers", workOffersHandlers.CreateWorkOffer)
	adminRoutes.GET("/work_offers/{id}/edit", workOffersHandlers.EditWorkOffer)
	adminRoutes.PUT("/work_offers/{id}", workOffersHandlers.UpdateWorkOffer)
	adminRoutes.DELETE("/work_offers/{id}", workOffersHandlers.DeleteWorkOffer)
}
