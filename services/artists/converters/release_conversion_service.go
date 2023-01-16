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
		"Description":        "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
		"ReleaseName":        release.Name,
		"ReleaseDate":        release.ReleaseDate.AsTime().Year(),
		"ImageDetails":       toImageDetailsMap(release.ImageDetails),
		"Tags":               []string{"indie", "post-punk"},
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
