package startup

import (
	"main/middlewares"
	startupModels "main/models/startup"
	templateFunctions "main/template-functions"
	"text/template"

	"github.com/gin-gonic/gin"
	consulClient "github.com/punk-link/consul-client"
	"github.com/punk-link/logger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func Configure(logger logger.Logger, consul consulClient.ConsulClient, appSecrets map[string]any, options *startupModels.StartupOptions) *gin.Engine {
	injector := buildDependencies(logger, consul, appSecrets)

	gin.SetMode(options.GinMode)
	app := gin.Default()

	notFoundMiddleware, _ := middlewares.NewNotFoundMiddleware(injector)
	app.NoRoute(notFoundMiddleware.HandlePageNotFound())

	app.Use(metricsMiddleware(options.ServiceName))
	app.Use(otelgin.Middleware(options.ServiceName))

	app.SetFuncMap(template.FuncMap{
		"convertToInt":        templateFunctions.ConvertToInt,
		"getPlatformIconPath": templateFunctions.GetPlatformIconPath,
		"getPlatformName":     templateFunctions.GetPlatformName,
		"sub":                 templateFunctions.Sub,
	})

	app.LoadHTMLGlob("./var/www/templates/**/*.go.tmpl")
	app.Static("/assets", "./var/www/assets")

	initSentry(app, logger, consul, options.EnvironmentName)
	configureOpenTelemetry(logger, consul, options)
	setupRouts(app, injector)

	return app
}
