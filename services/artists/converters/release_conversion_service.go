package converters

import (
	"fmt"
	commonServices "main/services/common"

	contracts "github.com/punk-link/presentation-contracts"
)

func ToSlimReleaseMaps(hashCoder commonServices.HashCoder, source []*contracts.SlimRelease) []map[string]any {
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

func ToReleaseMap(hashCoder commonServices.HashCoder, dataService commonServices.TemplateDataServer, release *contracts.Release) map[string]any {
	title := fmt.Sprintf("%s – %s", release.Name, release.ReleaseArtists[0].Name)
	tracks := toTrackMaps(hashCoder, release.Tracks)

	// TODO: add album and single numbers to ArtistStats
	return dataService.Enrich(title, map[string]any{
		"AlbumNumber":        int(release.ArtistStats.SoleReleaseNumber),
		"Artists":            ToSlimArtistMaps(hashCoder, release.ReleaseArtists),
		"CompilationNumber":  int(release.ArtistStats.CompilationNumber),
		"Date":               release.ReleaseDate.AsTime().Year(),
		"Description":        release.Description,
		"ImageDetails":       toImageDetailsMap(release.ImageDetails),
		"Name":               release.Name,
		"SingleNumber":       0,
		"Tags":               release.Tags,
		"Tracks":             tracks,
		"StreamingPlatforms": toPlatformUrlMaps(release.PlatformUrls),
	})
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

func toTrackMaps(hashCoder commonServices.HashCoder, tracks []*contracts.Track) []map[string]any {
	results := make([]map[string]any, len(tracks))
	for i, track := range tracks {
		results[i] = map[string]any{
			"Artists":    ToSlimArtistMaps(hashCoder, track.Artists),
			"IsExplicit": track.IsExplicit,
			"Name":       track.Name,
		}
	}

	return results
}
