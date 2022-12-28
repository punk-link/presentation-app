package startup

import (
	controllers "main/controllers"
	staticControllers "main/controllers/static"
	artistServices "main/services/artists"
	commonServices "main/services/common"
	releaseServices "main/services/releases"

	cacheManager "github.com/punk-link/cache-manager"
	consulClient "github.com/punk-link/consul-client"
	loggerService "github.com/punk-link/logger"
	"github.com/samber/do"
)

func buildDependencies(logger loggerService.Logger, consul consulClient.ConsulClient, appSecrets map[string]any) *do.Injector {
	injector := do.New()

	do.Provide(injector, func(i *do.Injector) (loggerService.Logger, error) {
		return loggerService.New(), nil
	})

	do.Provide(injector, registerCacheManager[map[string]any]())

	do.Provide(injector, artistServices.New)
	do.Provide(injector, commonServices.NewHashCoder)
	do.Provide(injector, releaseServices.New)

	do.Provide(injector, controllers.NewMetricsController)
	do.Provide(injector, controllers.NewStatusController)

	do.Provide(injector, staticControllers.NewArtistController)
	do.Provide(injector, staticControllers.NewReleaseController)

	return injector
}

func registerCacheManager[T any]() func(i *do.Injector) (cacheManager.CacheManager[T], error) {
	return func(i *do.Injector) (cacheManager.CacheManager[T], error) {
		return cacheManager.New[T](), nil
	}
}
