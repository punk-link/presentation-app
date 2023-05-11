package converters

import (
	"main/constants"
	commonServices "main/services/common"

	contracts "github.com/punk-link/presentation-contracts"
)

func ToArtistMap(hashCoder commonServices.HashCoder, dataService commonServices.TemplateDataServer, artist *contracts.Artist, sortedReleases []*contracts.SlimRelease, albumNumber int, singleNumber int, compilationNumber int) map[string]any {
	return dataService.Enrich(artist.Name, map[string]any{
		"Slug":              hashCoder.Encode(int(artist.Id)),
		"AlbumNumber":       albumNumber,
		"ArtistName":        artist.Name,
		"CompilationNumber": compilationNumber,
		"Releases":          ToSlimReleaseMaps(hashCoder, sortedReleases),
		"SingleNumber":      singleNumber,
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

// TODO: move to contracts
func toSocialNetworks(artist *contracts.Artist) map[string]any {
	result := make(map[string]any)

	result[constants.FACEBOOK] = "https://www.facebook.com/thesubways/"
	result[constants.INSTAGRAM] = "https://www.instagram.com/thesubways/"
	result[constants.TWITTER] = "https://twitter.com/thesubways"

	return result
}
