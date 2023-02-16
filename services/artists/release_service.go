package artists

import (
	"context"
	"main/helpers"
	"main/services/artists/converters"
	commonServices "main/services/common"
	"time"

	cacheManager "github.com/punk-link/cache-manager"
	"github.com/punk-link/logger"
	contracts "github.com/punk-link/presentation-contracts"
	"github.com/samber/do"
)

type ReleaseService struct {
	cache       cacheManager.CacheManager[map[string]any]
	dataService commonServices.TemplateDataServer
	grpcClient  contracts.PresentationClient
	hashCoder   commonServices.HashCoder
	logger      logger.Logger
}

func NewReleaseService(injector *do.Injector) (ReleaseServer, error) {
	cache := do.MustInvoke[cacheManager.CacheManager[map[string]any]](injector)
	dataService := do.MustInvoke[commonServices.TemplateDataServer](injector)
	grpcClient := do.MustInvoke[contracts.PresentationClient](injector)
	hashCoder := do.MustInvoke[commonServices.HashCoder](injector)
	logger := do.MustInvoke[logger.Logger](injector)

	return &ReleaseService{
		cache:       cache,
		dataService: dataService,
		grpcClient:  grpcClient,
		hashCoder:   hashCoder,
		logger:      logger,
	}, nil
}

func (t *ReleaseService) Get(hash string) (map[string]any, error) {
	cacheKey := helpers.BuildCacheKey("Release", hash)
	value, isCached := t.cache.TryGet(cacheKey)
	if isCached {
		return value, nil
	}

	id := t.hashCoder.Decode(hash)
	request := contracts.ReleaseRequest{Id: int32(id)}

	release, err := t.grpcClient.GetRelease(context.Background(), &request)
	result, err := buildReleaseResult(err, t.hashCoder, t.dataService, release)
	if err == nil {
		t.cache.Set(cacheKey, result, RELEASE_CACHE_DURATION)
	}

	return result, err
}

func buildReleaseResult(err error, hashCoder commonServices.HashCoder, dataService commonServices.TemplateDataServer, release *contracts.Release) (map[string]any, error) {
	if err != nil {
		return make(map[string]any), err
	}

	return converters.ToReleaseMap(hashCoder, dataService, release), nil
}

const RELEASE_CACHE_DURATION = time.Hour * 0
