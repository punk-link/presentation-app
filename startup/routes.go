package startup

import (
	"main/controllers"
	staticControllers "main/controllers/static"

	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func setupRouts(app *gin.Engine, injector *do.Injector) {
	registerRoutes(injector, func(controller *controllers.MetricsController) {
		app.GET("/metrics", controller.GetMetrics)
	})

	registerRoutes(injector, func(controller *controllers.StatusController) {
		app.GET("/health", controller.CheckHealth)
		app.GET("/", controller.Ok)
	})

	registerRoutes(injector, func(controller *controllers.HashController) {
		app.GET("/hashes/:target/decode", controller.Decode)
		app.GET("/hashes/:target/encode", controller.Encode)
	})

	registerRoutes(injector, func(controller *staticControllers.ArtistController) {
		app.GET("/artists/:hash", controller.Get)
		app.GET("/a/:hash", controller.Get)
	})

	registerRoutes(injector, func(controller *staticControllers.ReleaseController) {
		app.GET("/releases/:hash", controller.Get)
		app.GET("/r/:hash", controller.Get)
	})
}

func registerRoutes[T any](injector *do.Injector, function func(*T)) {
	controller := do.MustInvoke[*T](injector)
	function(controller)
}
