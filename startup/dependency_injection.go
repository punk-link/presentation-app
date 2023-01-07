package startup

import (
	"fmt"
	controllers "main/controllers"
	staticControllers "main/controllers/static"
	artistServices "main/services/artists"
	commonServices "main/services/common"

	contracts "github.com/punk-link/presentation-contracts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	cacheManager "github.com/punk-link/cache-manager"
	consulClient "github.com/punk-link/consul-client"
	loggerService "github.com/punk-link/logger"
	"github.com/samber/do"
)

func buildDependencies(logger loggerService.Logger, consul consulClient.ConsulClient, appSecrets map[string]any) *do.Injector {
	injector := do.New()

	grpcClient := getGrpcClient(logger, consul)
	do.ProvideValue(injector, grpcClient)

	do.Provide(injector, func(i *do.Injector) (loggerService.Logger, error) {
		return loggerService.New(), nil
	})

	do.Provide(injector, registerCacheManager[map[string]any]())

	do.Provide(injector, commonServices.NewHashCoder)
	do.Provide(injector, artistServices.NewArtistService)
	do.Provide(injector, artistServices.NewReleaseService)

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

func getGrpcClient(logger loggerService.Logger, consul consulClient.ConsulClient) contracts.PresentationClient {
	hostSettingsValue, err := consul.Get("GrpcSettings")
	if err != nil {
		logger.LogFatal(err, "Can't obtain gRPC host address: %s", err.Error())
	}

	hostSettings := hostSettingsValue.(map[string]any)
	grpcTarget := fmt.Sprintf("%s:%s", hostSettings["PresentationHost"].(string), hostSettings["PresentationPort"].(string))

	options := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(grpcTarget, options)
	if err != nil {
		logger.LogFatal(err, "gRPC connection error: %e", err.Error())
	}

	return contracts.NewPresentationClient(conn)
}
