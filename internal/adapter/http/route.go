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
	WarehouseController   *controller.WarehouseController
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
	warehouse := r.App.Group("/v1/warehouse")

	warehouse.GET("/collection-documents", r.WarehouseController.List)
	warehouse.GET("/collection-documents/", r.WarehouseController.List)

	warehouse.POST("/collection-documents", r.WarehouseController.Create)
	warehouse.POST("/collection-documents/", r.WarehouseController.Create)

	warehouse.GET("/collection-documents/:id", r.WarehouseController.Get)
	warehouse.GET("/collection-documents/:id/", r.WarehouseController.Get)

	warehouse.PUT("/collection-documents/:id", r.WarehouseController.Update)
	warehouse.PUT("/collection-documents/:id/", r.WarehouseController.Update)

	warehouse.DELETE("/collection-documents/:id", r.WarehouseController.Delete)
	warehouse.DELETE("/collection-documents/:id/", r.WarehouseController.Delete)
}
