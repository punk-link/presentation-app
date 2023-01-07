package artists

import (
	"context"
	"fmt"
	artistModels "main/models/artists"
	commonServices "main/services/common"
	"time"

	cacheManager "github.com/punk-link/cache-manager"
	"github.com/punk-link/logger"
	contracts "github.com/punk-link/presentation-contracts"
	"github.com/samber/do"
)

type ReleaseService struct {
	cache      cacheManager.CacheManager[map[string]any]
	grpcClient contracts.PresentationClient
	hashCoder  *commonServices.HashCoder
	logger     logger.Logger
}

func NewReleaseService(injector *do.Injector) (*ReleaseService, error) {
	cache := do.MustInvoke[cacheManager.CacheManager[map[string]any]](injector)
	grpcClient := do.MustInvoke[contracts.PresentationClient](injector)
	hashCoder := do.MustInvoke[*commonServices.HashCoder](injector)
	logger := do.MustInvoke[logger.Logger](injector)

	return &ReleaseService{
		cache:      cache,
		grpcClient: grpcClient,
		hashCoder:  hashCoder,
		logger:     logger,
	}, nil
}

func (t *ReleaseService) Get(hash string) (map[string]any, error) {
	cacheKey := buildReleaseCacheKey(hash)
	value, isCached := t.cache.TryGet(cacheKey)
	if isCached {
		return value, nil
	}

	id := t.hashCoder.Decode(hash)
	request := contracts.ReleaseRequest{Id: int32(id)}

	release, err := t.grpcClient.GetRelease(context.Background(), &request)
	tracks, err := t.buildTracks(err, release.Tracks, release.ReleaseArtists)
	//platformUrls, err := t.getPlatformReleaseUrls(err, id)
	result, err := buildReleaseResult(err, release, tracks)

	if err == nil {
		t.cache.Set(cacheKey, result, RELEASE_CACHE_DURATION)
	}

	return result, err
}

func (t *ReleaseService) buildTracks(err error, tracks []*contracts.Track, releaseArtists []*contracts.SlimArtist) ([]artistModels.SlimTrack, error) {
	if err != nil {
		return make([]artistModels.SlimTrack, 0), err
	}

	releaseArtistIds := make(map[int32]int, len(releaseArtists))
	for _, artist := range releaseArtists {
		releaseArtistIds[artist.Id] = 0
	}

	slimTracks := make([]artistModels.SlimTrack, len(tracks))
	for i, track := range tracks {
		trackArtists := make([]string, 0)
		for _, artist := range track.Artists {
			if _, isExist := releaseArtistIds[artist.Id]; !isExist {
				trackArtists = append(trackArtists, artist.Name)
			}
		}

		slimTracks[i] = artistModels.SlimTrack{
			ArtistNames: trackArtists,
			IsExplicit:  track.IsExplicit,
			Name:        track.Name,
		}
	}

	return slimTracks, err
}

func buildReleaseCacheKey(hash string) string {
	return fmt.Sprintf("ArtistStaticRelease::%s", hash)
}

func buildReleaseResult(err error, release *contracts.Release, tracks []artistModels.SlimTrack /*, platformUrls []platformModels.PlatformReleaseUrl*/) (map[string]any, error) {
	if err != nil {
		return make(map[string]any), err
	}

	return map[string]any{
		"PageTitle":    fmt.Sprintf("%s – %s", release.Name, release.ReleaseArtists[0].Name),
		"ArtistNames":  release.ReleaseArtists,
		"ReleaseName":  release.Name,
		"ReleaseDate":  release.ReleaseDate.AsTime().Year(),
		"ImageDetails": release.ImageDetails,
		"Tracks":       tracks,
		//"StreamingPlatforms": platformUrls,
	}, err
}

const RELEASE_CACHE_DURATION = time.Hour * 24
