package converters

import (
	commonServices "main/services/common"

	contracts "github.com/punk-link/presentation-contracts"
)

func ToArtistMap(hashCoder *commonServices.HashCoder, dataService *commonServices.TemplateDataService, artist *contracts.Artist, soleReleases []*contracts.SlimRelease, compilations []*contracts.SlimRelease) map[string]any {
	return dataService.Enrich(artist.Name, map[string]any{
		"ArtistName":        artist.Name,
		"SoleReleaseNumber": len(soleReleases),
		"CompilationNumber": len(compilations),
		"Releases":          ToSlimReleaseMaps(hashCoder, soleReleases),
		"Compilations":      ToSlimReleaseMaps(hashCoder, compilations),
	})
}

func ToSlimArtistMaps(hashCoder *commonServices.HashCoder, artists []*contracts.SlimArtist) []map[string]any {
	results := make([]map[string]any, len(artists))
	for i, artist := range artists {
		results[i] = toSlimArtistMap(hashCoder, artist)
	}

	return results
}

func toSlimArtistMap(hashCoder *commonServices.HashCoder, artist *contracts.SlimArtist) map[string]any {
	return map[string]any{
		"Slug": hashCoder.Encode(int(artist.Id)),
		"Name": artist.Name,
	}
}
