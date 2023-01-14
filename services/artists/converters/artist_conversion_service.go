package converters

import (
	commonServices "main/services/common"

	contracts "github.com/punk-link/presentation-contracts"
)

func ToArtistMap(hashCoder *commonServices.HashCoder, artist *contracts.Artist, soleReleases []*contracts.SlimRelease, compilations []*contracts.SlimRelease) map[string]any {
	return map[string]any{
		"PageTitle":         artist.Name,
		"ArtistName":        artist.Name,
		"SoleReleaseNumber": len(soleReleases),
		"CompilationNumber": len(compilations),
		"Releases":          ToSlimReleaseMaps(hashCoder, soleReleases),
		"Compilations":      ToSlimReleaseMaps(hashCoder, compilations),
	}
}
