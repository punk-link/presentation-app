package artists

import (
	"context"
	"fmt"
	"main/models/artists/enums"
	converters "main/services/artists/converters"
	commonServices "main/services/common"
	"sort"
	"time"

	cacheManager "github.com/punk-link/cache-manager"
	"github.com/punk-link/logger"
	contracts "github.com/punk-link/presentation-contracts"
	"github.com/samber/do"
)

type ArtistService struct {
	cache      cacheManager.CacheManager[map[string]any]
	grpcClient contracts.PresentationClient
	hashCoder  *commonServices.HashCoder
	logger     logger.Logger
}

func NewArtistService(injector *do.Injector) (*ArtistService, error) {
	cache := do.MustInvoke[cacheManager.CacheManager[map[string]any]](injector)
	grpcClient := do.MustInvoke[contracts.PresentationClient](injector)
	hashCoder := do.MustInvoke[*commonServices.HashCoder](injector)
	logger := do.MustInvoke[logger.Logger](injector)

	return &ArtistService{
		cache:      cache,
		grpcClient: grpcClient,
		hashCoder:  hashCoder,
		logger:     logger,
	}, nil
}

func (t *ArtistService) Get(hash string) (map[string]any, error) {
	cacheKey := buildArtistCacheKey(hash)
	value, isCached := t.cache.TryGet(cacheKey)
	if isCached {
		return value, nil
	}

	id := t.hashCoder.Decode(hash)
	request := contracts.ArtistRequest{Id: int32(id)}

	artist, err := t.grpcClient.GetArtist(context.Background(), &request)
	soleReleases, compilations, err := t.sortReleases(err, artist.Releases)
	result, err := buildArtistResult(err, t.hashCoder, artist, soleReleases, compilations)
	if err == nil {
		t.cache.Set(cacheKey, result, ARTIST_CACHE_DURATION)
	}

	return result, err
}

func (t *ArtistService) sortReleases(err error, releases []*contracts.Release) ([]*contracts.Release, []*contracts.Release, error) {
	if err != nil {
		return make([]*contracts.Release, 0), make([]*contracts.Release, 0), err
	}

	soleReleases := make([]*contracts.Release, 0)
	compilations := make([]*contracts.Release, 0)
	for _, release := range releases {
		if release.Type == enums.Compilation {
			compilations = append(compilations, release)
		} else {
			soleReleases = append(soleReleases, release)
		}
	}

	sortReleasesInternal(soleReleases)
	sortReleasesInternal(compilations)

	return soleReleases, compilations, err
}

func buildArtistCacheKey(hash string) string {
	return fmt.Sprintf("Artist::%s", hash)
}

func buildArtistResult(err error, hashCoder *commonServices.HashCoder, artist *contracts.Artist, soleReleases []*contracts.Release, compilations []*contracts.Release) (map[string]any, error) {
	if err != nil {
		return make(map[string]any, 0), err
	}

	return map[string]any{
		"PageTitle":         artist.Name,
		"ArtistName":        artist.Name,
		"SoleReleaseNumber": len(soleReleases),
		"CompilationNumber": len(compilations),
		"Releases":          converters.ToSlimRelease(hashCoder, append(soleReleases, compilations...)),
	}, err
}

func sortReleasesInternal(releases []*contracts.Release) {
	sort.SliceStable(releases, func(i, j int) bool {
		return releases[i].ReleaseDate.AsTime().After(releases[j].ReleaseDate.AsTime())
	})
}

const ARTIST_CACHE_DURATION = time.Hour * 24
