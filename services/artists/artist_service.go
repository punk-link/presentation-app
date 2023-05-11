package artists

import (
	"context"
	"main/helpers"
	converters "main/services/artists/converters"
	commonServices "main/services/common"
	"sort"
	"time"

	cacheManager "github.com/punk-link/cache-manager"
	"github.com/punk-link/logger"
	platformContracts "github.com/punk-link/platform-contracts"
	contracts "github.com/punk-link/presentation-contracts"
	"github.com/samber/do"
)

type ArtistService struct {
	cache       cacheManager.CacheManager[map[string]any]
	dataService commonServices.TemplateDataServer
	grpcClient  contracts.PresentationClient
	hashCoder   commonServices.HashCoder
	logger      logger.Logger
}

func NewArtistService(injector *do.Injector) (ArtistServer, error) {
	cache := do.MustInvoke[cacheManager.CacheManager[map[string]any]](injector)
	dataService := do.MustInvoke[commonServices.TemplateDataServer](injector)
	grpcClient := do.MustInvoke[contracts.PresentationClient](injector)
	hashCoder := do.MustInvoke[commonServices.HashCoder](injector)
	logger := do.MustInvoke[logger.Logger](injector)

	return &ArtistService{
		cache:       cache,
		dataService: dataService,
		grpcClient:  grpcClient,
		hashCoder:   hashCoder,
		logger:      logger,
	}, nil
}

func (t *ArtistService) Get(hash string) (map[string]any, error) {
	cacheKey := helpers.BuildCacheKey("Artist", hash)
	value, isCached := t.cache.TryGet(cacheKey)
	if isCached {
		return value, nil
	}

	id := t.hashCoder.Decode(hash)
	request := contracts.ArtistRequest{Id: int32(id)}

	artist, err := t.grpcClient.GetArtist(context.Background(), &request)
	result, err := t.buildArtistResult(err, t.hashCoder, t.dataService, artist, artist.Releases)
	if err == nil {
		t.cache.Set(cacheKey, result, ARTIST_CACHE_DURATION)
	}

	return result, err
}

func (t *ArtistService) buildArtistResult(err error, hashCoder commonServices.HashCoder, dataService commonServices.TemplateDataServer, artist *contracts.Artist, releases []*contracts.SlimRelease) (map[string]any, error) {
	sortedReleases, albumNumber, singleNumber, compilationNumber, err := sortReleasesAndGetNumberByType(err, releases)
	if err != nil {
		return make(map[string]any, 0), err
	}

	return converters.ToArtistMap(hashCoder, dataService, artist, sortedReleases, albumNumber, singleNumber, compilationNumber), nil
}

func sortReleasesAndGetNumberByType(err error, releases []*contracts.SlimRelease) ([]*contracts.SlimRelease, int, int, int, error) {
	if err != nil {
		return make([]*contracts.SlimRelease, 0), 0, 0, 0, err
	}

	albumNumber := 0
	singleNumber := 0

	soleReleases := make([]*contracts.SlimRelease, 0)
	compilations := make([]*contracts.SlimRelease, 0)
	for _, release := range releases {
		if release.Type == platformContracts.Compilation {
			compilations = append(compilations, release)
		} else {
			if release.Type == platformContracts.Album {
				albumNumber++
			} else {
				singleNumber++
			}

			soleReleases = append(soleReleases, release)
		}
	}

	sortReleasesInternal(soleReleases)
	sortReleasesInternal(compilations)

	return append(soleReleases, compilations...), albumNumber, singleNumber, len(compilations), err
}

func sortReleasesInternal(releases []*contracts.SlimRelease) {
	sort.SliceStable(releases, func(i, j int) bool {
		return releases[i].ReleaseDate.AsTime().After(releases[j].ReleaseDate.AsTime())
	})
}

const ARTIST_CACHE_DURATION = time.Hour * 24
