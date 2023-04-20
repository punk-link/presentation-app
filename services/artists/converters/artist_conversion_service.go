package converters

import (
	commonServices "main/services/common"

	contracts "github.com/punk-link/presentation-contracts"
)

func ToArtistMap(hashCoder commonServices.HashCoder, dataService commonServices.TemplateDataServer, artist *contracts.Artist, soleReleases []*contracts.SlimRelease, compilations []*contracts.SlimRelease) map[string]any {
	return dataService.Enrich(artist.Name, map[string]any{
		"ArtistName":        artist.Name,
		"SoleReleaseNumber": int32(len(soleReleases)),
		"CompilationNumber": int32(len(compilations)),
		"Releases":          ToSlimReleaseMaps(hashCoder, soleReleases),
		"Compilations":      ToSlimReleaseMaps(hashCoder, compilations),
		"SocialNetworks":    toSocialNetworks(artist),
	})
}

func ToSlimArtistMaps(hashCoder commonServices.HashCoder, artists []*contracts.SlimArtist) []map[string]any {
	results := make([]map[string]any, len(artists))
	for i, artist := range artists {
		results[i] = toSlimArtistMap(hashCoder, artist)
	}

	return results
}

func toSlimArtistMap(hashCoder commonServices.HashCoder, artist *contracts.SlimArtist) map[string]any {
	return map[string]any{
		"Slug": hashCoder.Encode(int(artist.Id)),
		"Name": artist.Name,
	}
}

func toSocialNetworks(artist *contracts.Artist) map[string]any {
	result := make(map[string]any)

	result["facebook"] = "https://www.facebook.com/thesubways/"
	result["instagram"] = "https://www.instagram.com/thesubways/"
	result["twitter"] = "https://twitter.com/thesubways"

	return result
}
