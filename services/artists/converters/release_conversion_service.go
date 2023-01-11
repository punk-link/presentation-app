package converters

import (
	artistModels "main/models/artists"
	commonServices "main/services/common"

	contracts "github.com/punk-link/presentation-contracts"
)

func ToSlimRelease(hashCoder *commonServices.HashCoder, source []*contracts.SlimRelease) []artistModels.SlimRelease {
	results := make([]artistModels.SlimRelease, len(source))
	for i, release := range source {
		results[i] = artistModels.SlimRelease{
			Slug:         hashCoder.Encode(int(release.Id)),
			ImageDetails: release.ImageDetails,
			Name:         release.Name,
			ReleaseDate:  release.ReleaseDate.AsTime(), //.Year()
			Type:         release.Type,
		}
	}
	return results
}
