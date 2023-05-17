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
	tracks := toTrackMaps(hashCoder, release.Tracks, release.ReleaseArtists)

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

func getArtistPresentedOnAllTracksMap(artists []*contracts.SlimArtist, tracks []*contracts.Track) map[string]bool {
	presentedOnAllTracksArtistMap := make(map[string]bool)
	for _, releaseArtist := range artists {
		isPresentedOnAllTracks := false
		for _, track := range tracks {
			for _, trackArtist := range track.Artists {
				if trackArtist.Name == releaseArtist.Name {
					isPresentedOnAllTracks = true
					break
				}
			}

			if !isPresentedOnAllTracks {
				break
			}
		}

		if isPresentedOnAllTracks {
			presentedOnAllTracksArtistMap[releaseArtist.Name] = true
		}
	}

	return presentedOnAllTracksArtistMap
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

func toTrackMaps(hashCoder commonServices.HashCoder, tracks []*contracts.Track, artists []*contracts.SlimArtist) []map[string]any {
	presentedOnAllTracksArtistMap := getArtistPresentedOnAllTracksMap(artists, tracks)

	results := make([]map[string]any, len(tracks))
	for i, track := range tracks {
		trackArtists := make([]*contracts.SlimArtist, 0)
		for _, trackArtist := range track.Artists {
			if !presentedOnAllTracksArtistMap[trackArtist.Name] {
				trackArtists = append(trackArtists, trackArtist)
			}
		}

		results[i] = map[string]any{
			"Artists":    ToSlimArtistMaps(hashCoder, trackArtists),
			"IsExplicit": track.IsExplicit,
			"Name":       track.Name,
		}
	}

	return results
}
