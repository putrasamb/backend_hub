package httproute

import (
	"backend_hub/internal/adapter/http/controller"

	"github.com/labstack/echo/v4"
)

type Route struct {
	App                   *echo.Echo
	HealthCheckController *controller.HealthCheckController
	LogController         *controller.LogController
	TesController         *controller.TesController
	CollectionDocument    *controller.CollectionDocumentController
}

func (r *Route) Setup() {
	r.App.GET("/health", r.HealthCheckController.Ping)
	r.App.GET("/logs/:year/:month/:day/:f", r.LogController.Read)

	r.setupTes()
	r.SetupCollectionDocument()
}

func (r *Route) setupTes() {
	tes := r.App.Group("/tes")

	tes.GET("/v1", r.TesController.List)
	tes.GET("/v1/", r.TesController.List)
}

func (r *Route) SetupCollectionDocument() {
	collection := r.App.Group("/v1")

	collection.GET("/collection-documents", r.CollectionDocument.List)
	collection.GET("/collection-documents/", r.CollectionDocument.List)
}
