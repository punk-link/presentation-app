package converters

import (
	commonServices "main/services/common"

	contracts "github.com/punk-link/presentation-contracts"
)

func ToArtistMap(hashCoder commonServices.HashCoder, dataService commonServices.TemplateDataServer, artist *contracts.Artist, sortedReleases []*contracts.SlimRelease) map[string]any {
	return dataService.Enrich(artist.Name, map[string]any{
		//"Slug":               hashCoder.Encode(int(artist.Id)),
		"Artists":     toSlimArtistMaps(hashCoder, artist),
		"AlbumNumber": int(artist.ReleaseStats.AlbumNumber),
		//"Name":               artist.Name,
		"CompilationNumber":  int(artist.ReleaseStats.CompilationNumber),
		"PresentationConfig": artist.PresentationConfig,
		"Releases":           ToSlimReleaseMaps(hashCoder, sortedReleases),
		"SingleNumber":       int(artist.ReleaseStats.SingleNumber),
		"SocialNetworks":     artist.SocialNetworks,
	})
}

func ToSlimArtistMaps(hashCoder commonServices.HashCoder, artists []*contracts.SlimArtist) []map[string]any {
	results := make([]map[string]any, len(artists))
	for i, artist := range artists {
		results[i] = toSlimArtistMap(hashCoder, artist.Id, artist.Name)
	}

	return results
}

func toSlimArtistMaps(hashCoder commonServices.HashCoder, artist *contracts.Artist) []map[string]any {
	slimArtistMap := toSlimArtistMap(hashCoder, artist.Id, artist.Name)

	return []map[string]any{slimArtistMap}
}

func toSlimArtistMap(hashCoder commonServices.HashCoder, id int32, name string) map[string]any {
	return map[string]any{
		"Slug": hashCoder.Encode(int(id)),
		"Name": name,
	}
}
