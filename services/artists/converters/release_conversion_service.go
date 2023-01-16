package converters

import (
	"fmt"
	commonServices "main/services/common"

	contracts "github.com/punk-link/presentation-contracts"
)

func ToSlimReleaseMaps(hashCoder *commonServices.HashCoder, source []*contracts.SlimRelease) []map[string]any {
	results := make([]map[string]any, len(source))
	for i, release := range source {
		results[i] = map[string]any{
			"Slug":         hashCoder.Encode(int(release.Id)),
			"ImageDetails": toImageDetailsMap(release.ImageDetails),
			"Name":         release.Name,
			"ReleaseDate":  release.ReleaseDate.AsTime().Year(),
			"Type":         release.Type,
		}
	}

	return results
}

func ToReleaseMap(release *contracts.Release) map[string]any {
	title := fmt.Sprintf("%s – %s", release.Name, release.ReleaseArtists[0].Name)

	releaseArtistIds := make(map[int32]int, len(release.ReleaseArtists))
	for _, artist := range release.ReleaseArtists {
		releaseArtistIds[artist.Id] = 0
	}

	tracks := toTrackMaps(release.Tracks, releaseArtistIds)

	return map[string]any{
		"PageTitle":          title,
		"ArtistNames":        release.ReleaseArtists,
		"Description":        release.Description,
		"ReleaseName":        release.Name,
		"ReleaseDate":        release.ReleaseDate.AsTime().Year(),
		"ImageDetails":       toImageDetailsMap(release.ImageDetails),
		"Tags":               release.Tags,
		"Tracks":             tracks,
		"StreamingPlatforms": toPlatformUrlMaps(release.PlatformUrls),
	}
}

func toPlatformUrlMaps(platformUrls []*contracts.PlatformUrl) []map[string]any {
	results := make([]map[string]any, len(platformUrls))
	for i, url := range platformUrls {
		results[i] = map[string]any{
			"Id":  url.PlatformId,
			"Url": url.Url,
		}
	}

	return results
}

func toTrackMaps(tracks []*contracts.Track, artistIds map[int32]int) []map[string]any {
	results := make([]map[string]any, len(tracks))
	for i, track := range tracks {
		artistNames := make([]string, 0)
		for _, artist := range track.Artists {
			if _, isExist := artistIds[artist.Id]; !isExist {
				artistNames = append(artistNames, artist.Name)
			}
		}

		results[i] = map[string]any{
			"ArtistNames": artistNames,
			"IsExplicit":  track.IsExplicit,
			"Name":        track.Name,
		}
	}

	return results
}
