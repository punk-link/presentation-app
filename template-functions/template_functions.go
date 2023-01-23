package templateFunctions

import platformContracts "github.com/punk-link/platform-contracts"

func GetPlatformIconPath(platformId string) string {
	switch platformId {
	case platformContracts.Spotify:
		return "Spotyfy.webp"
	case platformContracts.Deezer:
		return "Deezer.webp"
	}

	return ""
}

func GetPlatformName(platformId string) string {
	switch platformId {
	case platformContracts.Spotify:
		return "Spotyfy"
	case platformContracts.Deezer:
		return "Deezer"
	}

	return ""
}
